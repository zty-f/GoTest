package http

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// HttpResponse 请求结果数据结构
type HttpResponse struct {
	http.Response
	URL  string
	Body []byte
}

// NewResponse 默认返回数据结构
func NewResponse() *HttpResponse {
	return &HttpResponse{Body: []byte("{}")}
}

// 定义并初始化客户端变量
var client *http.Client

func getClinet() *http.Client {
	if client == nil {
		client = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // 信任所有证书
				Dial: func(netw, addr string) (net.Conn, error) {
					c, err := net.DialTimeout(netw, addr, time.Second*2) // 限制建立TCP连接的时间
					if err != nil {
						return nil, err
					}
					return c, nil

				},
				MaxIdleConnsPerHost:   10,
				ResponseHeaderTimeout: time.Second * 3, // 限制读取response header的时间
			},
			Timeout: 5 * time.Second, // 从连接(Dial)到读完response body 的时间
		}
	}

	return client
}

func SimpleClient(urls string, method string, header map[string]string, params interface{}) (*HttpResponse, error) {
	// Service.Flagtime("1")
	var pbody io.Reader
	req, err := http.NewRequest(method, urls, nil)
	if err != nil {
		return nil, err
	}

	if params != nil {
		if strings.ToUpper(method) == "GET" {
			if post, ok := params.(map[string]string); ok {
				q := req.URL.Query()
				for k, v := range post {
					q.Add(k, v)
				}
				req.URL.RawQuery = q.Encode()
			}

			// Common.SetDebug(fmt.Sprintf("Send HTTP Query: %s", urls+"?"+req.URL.RawQuery), 2)

		} else if strings.ToUpper(method) == "POST" {
			if post, ok := params.(map[string]string); ok {
				data := make(url.Values)
				for k, v := range post {
					data.Add(k, string(v))
				}
				pbody = strings.NewReader(data.Encode())
			}

			if post, ok := params.([]byte); ok {
				pbody = bytes.NewReader(post)
			}

			if req, err = http.NewRequest(method, urls, pbody); err != nil {
				// logcus.Error("CacheHTTP gen newRequest:" + err.Error())
				return nil, err
			}
			// req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			// Common.SetDebug(fmt.Sprintf("Send HTTP Query: %s", urls), 2)
		}
	}

	//增加header
	req.Header.Add("User-Agent", "Mozilla/5.0")

	for k, v := range header {
		req.Header.Set(k, v)
	}

	httpRes := NewResponse()
	client := getClinet()
	resp, err := client.Do(req)
	if err != nil {
		// Common.SetDebug(fmt.Sprintf("HTTP Query Downgrade: %s", err.Error()), 2)
		// logcus.Error("CacheHTTP request error: " + err.Error())

		return nil, err
	}

	// if resp.StatusCode != 200 {
	// 	//不抛出错误而是接口降级
	// 	// Common.SetDebug(fmt.Sprintf("HTTP Query Downgrade: non-200 StatusCode:%s", urls), 2)
	// 	logcus.Error("CacheHTTP request got non-200 StatusCode: " + urls)

	// 	httpRes.HttpStatus = resp.Status
	// 	httpRes.HttpStatusCode = resp.StatusCode
	// 	return httpRes
	// }

	// Common.SetDebug(fmt.Sprintf("HTTP Query Result{"+Service.Flagtime("1")+"} : status :%s, content length:%d, url:%s", resp.Status, resp.ContentLength, urls), 2)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("decode resp.Body error: %s", err.Error())
	}

	httpRes.URL = urls
	httpRes.Body = body
	httpRes.Response = *resp

	// httpRes.HttpStatus = resp.Status
	// httpRes.HttpStatusCode = resp.StatusCode
	// httpRes.ContentLength = resp.ContentLength

	return httpRes, nil
}
