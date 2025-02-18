package http

import (
	"context"
	"errors"
	"fmt"
	"github.com/bytedance/sonic"
	"net/http"
	"test/http_utils"
	"test/util"
	"testing"
)

type Req struct {
}

type Resp struct {
	Stat    int    `json:"stat"` // 结果状态标识，1成功，其他失败
	Msg     string `json:"msg"`
	Data    Data   `json:"data"`
	TraceId string `json:"trace_id"`
}

type Data struct {
}

const API = "https://xxxx.xxx.com/xx/xxx"

func doRequestOfProxy(ctx context.Context, req Req) (Data, error) {
	var data Data
	get := http_utils.Map2HttpQueryEncode(http_utils.Struct2ReqMap(req))
	// post := http_utils.Map2HttpQueryEncode(http_utils.Struct2ReqMap(req))  //post一致
	p, err := NewProxy(API, get, "", http.MethodGet, 3, map[string]string{})
	if err != nil {
		return data, err
	}
	res, err := p.Do(ctx)
	if err != nil {
		return data, err
	}
	if res.Code > 0 {
		return data, fmt.Errorf("proxy return error code[%d], msg[%s]", res.Code, res.Message)
	}
	err = util.CheckStringForUnmarshal(res.Data.DecodeRes)
	if err != nil {
		if errors.Is(err, util.EmptyErrMsg) {
			return data, nil
		}
		return data, fmt.Errorf("proxy return error msg[%s]", err)
	}

	apiResp := Resp{}
	err = sonic.UnmarshalString(res.Data.DecodeRes, &apiResp)
	if err != nil {
		return data, fmt.Errorf("proxy return error msg[%s]", err)
	}
	if apiResp.Stat != 1 {
		return data, fmt.Errorf("proxy return error msg[%s]", apiResp.Msg)
	}
	return apiResp.Data, nil
}

func TestNewProxy(t *testing.T) {
	data, err := doRequestOfProxy(context.Background(), Req{})
	if err != nil {
		return
	}
	fmt.Println(data)
}
