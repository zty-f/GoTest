package http

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/cast"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"test/http_utils"
	"test/util"
	"time"
)

var proxyConf map[string]interface{}

const PROXY_API = "http://xxx.xxx.com/proxyapi/v1/proxy/index"

// 代理参数
type Proxy struct {
	// 需要代理的接口
	Url string `json:"url"`
	// 需要代理的接口 get 参数
	Get string `json:"get"`
	// 需要代理的接口 post 参数
	Post string `json:"post"`
	// 需要代理的接口超时时间，不传递的情况下默认为 3
	Timeout int `json:"timeout"`
	// 需要代理的服务请求方法
	Method string `json:"method"`
	// 特殊 header ，注: 美校网关鉴权参数由代理服务自动处理，这里仅用来指定除了网关鉴权以外地其他 header，例如cookie
	Headers map[string]string `json:"headers"`
}

// 代理返回结果
type proxyRes struct {
	// 代理服务错误码，0 为 ok，非 0 为错误
	Code int `json:"code"`
	// 代理服务错误信息
	Message string `json:"message"`
	// 代理接口的返回值
	// 特别注意：各个服务的返回值可能不一样，这里原样返回，业务侧自行处理 DecodeRes 即可
	Data struct {
		EncodeRes string `json:"encode_res"`
		DecodeRes string `json:"decode_res"`
	}
}

// NewProxy 生成代理对象
func NewProxy(api, get, post, method string, timeout int, headers map[string]string) (*Proxy, error) {
	if _, err := url.ParseRequestURI(api); err != nil {
		return nil, err
	}

	if len(method) < 3 {
		method = http.MethodGet
	}
	method = strings.ToUpper(method)
	if method != http.MethodPost && method != http.MethodGet {
		return nil, errors.New("method is invalid")
	}

	if timeout < 1 {
		timeout = 3
	}

	return &Proxy{
		Url:     api,
		Get:     get,
		Post:    post,
		Timeout: timeout,
		Method:  method,
		Headers: headers,
	}, nil
}

// Do 发起请求
func (p *Proxy) Do(c context.Context) (*proxyRes, error) {
	proxyRes := &proxyRes{
		Code:    1,
		Message: "Request failed",
	}

	u, err := url.Parse(p.Url)
	if err != nil {
		return proxyRes, err
	}

	fmt.Println(u)

	// 如果是指定域名的服务直接切走，替换域名直接请求，不走代理
	// if isSyhService, _ := SyhService[u.Host]; isSyhService {
	//	p.Url = strings.ReplaceAll(p.Url, "xxx.xxx", "xxx.xxx")
	//	return p.DoDirectRequest(c)
	// }

	// 测试环境默认走代理
	// if utils.IsDevEnv() {
	//	return p.DoProxyRequest(c)
	// }

	// 内网域名直接请求，不走代理
	// isIntranetNet, _ := IntranetRequestHostMap[u.Host]
	// if isIntranetNet {
	//	return p.DoDirectRequest(c)
	// }
	return p.DoProxyRequest(c)
}

// DoProxyRequest 发起代理请求
func (p *Proxy) DoProxyRequest(c context.Context) (*proxyRes, error) {

	proxyResp := &proxyRes{
		Code:    1,
		Message: "Request failed",
	}
	var err error

	params, err := p.makeProxySign(proxyConf)
	if err != nil {
		return proxyResp, err
	}

	dialer := &net.Dialer{
		Timeout: time.Duration(int(proxyConf["timeout"].(int)) * int(time.Second)),
	}
	http.DefaultTransport.(*http.Transport).DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
		// 开发环境和测试环境单独处理代理配置
		// if utils.IsDevEnv() || utils.IsTestEnv() {
		//	resolveHttps := proxyConf["resolve_https"].(string)
		//	resolveHttp := proxyConf["resolve_http"].(string)
		//	if addr == "xxx.xxx.com:443" && len(resolveHttps) > 1 {
		//		addr = resolveHttps
		//	}
		//	if addr == "xxx.xxx.com:80" && len(resolveHttp) > 1 {
		//		addr = resolveHttp
		//	}
		// }
		return dialer.DialContext(ctx, network, addr)
	}
	resp, err := http.Post(PROXY_API, "application/json", params)
	if err != nil {
		return proxyResp, err
	}
	if resp.StatusCode != http.StatusOK {
		return proxyResp, errors.New("proxyapi response status is " + cast.ToString(resp.StatusCode))
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return proxyResp, err
	}
	err = json.Unmarshal(respBytes, &proxyResp)
	if err != nil {
		return proxyResp, err
	}

	if len(proxyResp.Data.EncodeRes) > 1 {
		decodeRes, err := util.CbcDecode(proxyResp.Data.EncodeRes, proxyConf["app_key"].(string), proxyConf["app_iv"].(string))
		if err != nil {
			return proxyResp, err
		}
		proxyResp.Data.DecodeRes = decodeRes
	}

	return proxyResp, nil
}

// DoDirectRequest 发起直接请求
func (p *Proxy) DoDirectRequest(c context.Context) (*proxyRes, error) {
	proxyResp := &proxyRes{
		Code:    1,
		Message: "Request failed",
	}
	var err error

	url := p.Url
	if p.Method == http.MethodGet && p.Get != "" {
		url += "?" + p.Get
	}
	data := strings.NewReader("")
	if p.Method == http.MethodPost {
		data = strings.NewReader(p.Post)
	}
	ct, ok := p.Headers["Content-Type"]
	if p.Method == http.MethodPost && (len(ct) < 1 || !ok) {
		p.Headers["Content-Type"] = "application/x-www-form-urlencoded; charset=utf-8"
	}

	sourceTraceId := c.Value("x_trace_id")
	traceIdStr, ok := sourceTraceId.(string)
	if ok {
		p.Headers["traceid"] = traceIdStr
	}

	pkg := http_utils.CallPkg{
		Ctx:         c,
		Timeout:     p.Timeout,
		Uri:         url,
		Method:      p.Method,
		Header:      p.Headers,
		ContentType: "",
		Data:        data,
	}
	respBytes, err := http_utils.CallHttp(pkg)
	if err != nil {
		return proxyResp, err
	}

	proxyResp = &proxyRes{
		Code:    0,
		Message: "succ",
		Data: struct {
			EncodeRes string `json:"encode_res"`
			DecodeRes string `json:"decode_res"`
		}{
			EncodeRes: "",
			DecodeRes: string(respBytes),
		},
	}

	return proxyResp, nil
}

// makeProxySign 生成请求+签名
func (p *Proxy) makeProxySign(proxyConf map[string]interface{}) (*strings.Reader, error) {
	params := make(map[string]interface{})
	params["app_id"] = proxyConf["app_id"]
	params["timestamp"] = time.Now().Unix()
	body, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	params["body"] = string(body)
	params["sign"] = Md5X(HttpBuildQuery(params) + proxyConf["app_key"].(string))

	return strings.NewReader(util.Strval(params)), nil
}

// HttpBuildQuery 生成 http query
func HttpBuildQuery(data map[string]interface{}) string {
	var uri url.URL
	q := uri.Query()
	for k, v := range data {
		q.Add(k, util.Strval(v))
	}
	return q.Encode()
}

// md5
func Md5X(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// 转换struct为map[string]interface{}
// struct的每个元素的json tag 作为map的key，value作为map的value
// TODO 目前只能转一层嵌套，多层的会有问题
// TODO 注意下面的过滤情况，尤其是没有json的tag， 以及number类型的value， 如果将int=0 —> string="0"则不会过滤
/* struct中这些元素不会被放到map中
1.没有json tag的元素
2.kind==reflect.string && value==""
3.kind==reflect.Slice  && (value==nil || len(value)==0)
4.kind==reflect.Struct && value==nil
5.kind==reflect.Int | reflect.Float 等数值类型 && value==0 | 0.0 等默认值
*/
func Struct2ReqMap(data interface{}) map[string]interface{} {
	res := map[string]interface{}{}
	if data == nil {
		return res
	}

	v := reflect.TypeOf(data)
	reflectValue := reflect.ValueOf(data)
	reflectValue = reflect.Indirect(reflectValue)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		tag := v.Field(i).Tag.Get("json")
		field := reflectValue.Field(i)

		if tag == "" || tag == "-" {
			continue
		}

		if v.Field(i).Type.Kind() == reflect.Ptr {
			field = field.Elem()
		}

		if !field.IsValid() || field.IsZero() {
			continue
		}

		switch v.Field(i).Type.Kind() {
		case reflect.Struct:
			res[tag] = Struct2ReqMap(field.Interface())
		case reflect.Slice:
			if field.Len() > 0 {
				res[tag] = field.Interface()
			}
		default:
			res[tag] = field.Interface()
		}

	}
	return res
}

func Map2HttpQueryEncode(m map[string]interface{}) string {
	return Map2HttpQuery(m).Encode()
}

// 转换map[string]interface{} 为url.Values
// TODO 目前不支持map的value=interface{}是复杂结构的情况， value=slice可以支持一层基础类型的转换，如 []int  []string等
// 这个函数主要是提供给 Struct2ReqMap 函数的返回值使用的, 如果直接构造了map[string]interface{}，也是可以使用的
// 两个函数加起来的主要工作是将 struct结构体 转换为 url.Values 的 HttpQuery/encodeUrl的键值对参数形式
// 通过 url.Values.Encode() 方法生成string类型的参数对， 如果想直接获得string，使用 Map2HttpQueryEncode
// 返回的 url.Values, 可以通过 url.Values.Add(key, value)来增加参数对，再进行url.Values.Encode()
// 将string拼接到Get请求的url后面， 或者作为 Content-Type=x-www-form-urlencoded 的post请求的的body
// example:
//
//	m := map[string]interface{}{"age":20, "name":"Tom"}
//	Map2HttpQuery(m).Encode()   -->  age=20&name=Tom
func Map2HttpQuery(m map[string]interface{}) url.Values {
	var uri = url.Values{}
	if m == nil {
		return uri
	}

	for k, v := range m {
		vv := reflect.ValueOf(v)
		if vv.Kind() == reflect.Ptr {
			vv = vv.Elem()
		}

		if !vv.IsValid() {
			continue
		}

		switch vv.Kind() {
		case reflect.Slice:
			if vv.IsNil() {
				continue
			}
			if vv.Len() < 0 {
				continue
			}

			length := vv.Len()
			for i := 0; i < length; i++ {
				element := vv.Index(i)
				uri.Add(fmt.Sprintf("%s[]", k), ToString(element.Interface()))
			}
		case reflect.Map, reflect.Struct:
			byteJson, _ := json.Marshal(v)
			uri.Add(k, string(byteJson))
		default:
			uri.Add(k, ToString(v))
		}

	}
	return uri
}

// 调用cast.ToStringE方法
// 该方法对结构体和slice等不做处理
// 这里捕获err!=nil，做一个json序列化
func ToString(ifc interface{}) string {
	ret, err := cast.ToStringE(ifc)
	if err != nil {
		jsonByte, err := json.Marshal(ret)
		if err != nil {
			return ""
		}
		return string(jsonByte)
	}

	return ret
}
