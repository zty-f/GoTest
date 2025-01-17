package http_utils

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"net/url"
	"reflect"
)

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
