package http_utils

import (
	"context"
	"fmt"
	"github.com/valyala/fasthttp"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type CallPkg struct {
	Ctx         context.Context
	Timeout     int
	Uri         string
	Method      string
	Header      map[string]string
	ContentType string
	Data        io.Reader
}

var client *fasthttp.Client

func init() {
	readTimeout, _ := time.ParseDuration("5000ms")
	writeTimeout, _ := time.ParseDuration("5000ms")
	client = &fasthttp.Client{
		ReadTimeout:                   readTimeout,
		WriteTimeout:                  writeTimeout,
		MaxIdleConnDuration:           10 * time.Second, // 空闲链接时间应短，避免请求服务的 keep-alive 过短主动关闭，默认10秒
		NoDefaultUserAgentHeader:      true,             // Don't send: User-Agent: fasthttp
		DisableHeaderNamesNormalizing: true,             // If you set the case on your headers correctly you can enable this
		DisablePathNormalizing:        true,
		// increase DNS cache time to an hour instead of default minute
		Dial: func(addr string) (net.Conn, error) {
			return fasthttp.DialTimeout(addr, time.Duration(60)*time.Second)
		},
	}
}

// 调用http的简单封装
//
//	入参为 CallPkg {
//		Ctx          context.Context
//	 Timeout      超时 单位毫秒
//		Uri          http地址资源
//	 Method       目前只支持GET POST
//	 Header       请求头参数
//	 ContentType  目前针对的是post请求 使用form或者json。 application/x-www-form-urlencoded   application/json
//	 Data         io.Reader
//	}
func CallHttp(pkg CallPkg) (respData []byte, err error) {
	// 检验uri
	if _, err := url.ParseRequestURI(pkg.Uri); err != nil {
		return respData, err
	}

	// 检验方法
	pkg.Method = strings.ToUpper(pkg.Method)
	if len(pkg.Method) < 3 {
		pkg.Method = http.MethodGet
	}

	// 检验timeout
	if pkg.Timeout < 0 {
		pkg.Timeout = 0
	}

	// 检验context
	if pkg.Ctx == nil {
		pkg.Ctx = context.Background()
	}
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseResponse(resp)
		fasthttp.ReleaseRequest(req)
	}()
	req.SetRequestURI(pkg.Uri)

	// 设置header
	for k, v := range pkg.Header {
		req.Header.Set(k, v)
	}

	if pkg.Method == http.MethodPost && pkg.ContentType != "" {
		req.Header.Set("Content-Type", pkg.ContentType)
	}

	timeout := time.Duration(pkg.Timeout) * time.Millisecond
	req.SetTimeout(timeout)
	switch pkg.Method {
	case http.MethodGet:
		req.Header.SetMethod(http.MethodGet)
	case http.MethodPost:
		bytes, err := io.ReadAll(pkg.Data)
		if err != nil {
			return respData, err
		}
		req.SetBody(bytes)
		req.Header.SetMethod(http.MethodPost)
	default:
		if pkg.Method != http.MethodGet && pkg.Method != http.MethodPost {
			return respData, fmt.Errorf("invalid or not support method:%s", pkg.Method)
		}
	}
	err = client.Do(req, resp)
	if err != nil {
		return respData, err
	}
	respBody := make([]byte, len(resp.Body()))
	copy(respBody, resp.Body())
	return respBody, err
}
