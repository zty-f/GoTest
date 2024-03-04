package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"
)

type Response struct {
	Body []byte
}

func TestSimpleClient(t *testing.T) {
	request, err := doRequestOfSimpleClient()
	if err != nil {
		return
	}
	fmt.Println(string(request.Body))
}

func doRequestOfSimpleClient() (res *Response, err error) {
	url := ""
	header := map[string]string{}
	params := map[string]string{}
	jsonParams, _ := json.Marshal(params)
	ret, err := SimpleClient(url, "GET", header, jsonParams)
	if err != nil {
		return
	}
	if ret.StatusCode != 200 {
		err = errors.New("接口调用失败")
		return
	}
	err = json.Unmarshal(ret.Body, &res)
	if err != nil {
		return
	}
	return
}
