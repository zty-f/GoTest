package resp

import (
	"encoding/json"
	"net/http"
	"test/errors"

	"github.com/gin-gonic/gin"
)

const TraceIDKey = "x_trace_id"

type ReJSON struct {
	Stat    errors.Stat `json:"stat"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Msg     string      `json:"msg"` // 作用和Message作用一样，兼容客户端统一使用msg字段
	Data    interface{} `json:"data"`
	TraceID string      `json:"trace_id"`
}

type ReErrJSON struct {
	ReJSON
	Metadata map[string]string `json:"metadata"`
}

var emptyData map[string]interface{}

// WriteSuccessJSON write success response with struct
func WriteSuccessJSON(c *gin.Context, data interface{}) {
	if data == nil {
		data = emptyData
	}

	res := ReJSON{
		Stat:    errors.Stat_SUCCESS,
		Code:    errors.SuccessCode,
		Message: errors.SuccessMessage,
		Msg:     errors.SuccessMessage,
		Data:    data,
		TraceID: c.GetString(TraceIDKey),
	}

	resByte, _ := json.Marshal(res)
	c.Set("output", string(resByte))

	c.JSON(http.StatusOK, res)
}

// WriteErrJSON write response when biz return err
func WriteErrJSON(c *gin.Context, err error) {
	if err != nil {
		se := errors.FromError(err)

		res := ReErrJSON{
			ReJSON: ReJSON{
				Stat:    se.Stat,
				Code:    int(se.Code),
				Message: se.Message,
				Msg:     se.Message,
				Data:    emptyData,
				TraceID: c.GetString(TraceIDKey),
			},
			Metadata: se.Metadata,
		}
		resByte, _ := json.Marshal(res)
		c.Set("output", string(resByte))
		c.JSON(http.StatusOK, res)

		return
	}

	resByte, _ := json.Marshal(emptyData)
	c.Set("output", string(resByte))

	WriteSuccessJSON(c, emptyData)
}

// WriteErrJSONWithData write response when biz return err
// don't recommend use because the params
func WriteErrJSONWithData(c *gin.Context, err error, data interface{}) {
	if data == nil {
		data = emptyData
	}

	if err != nil {
		se := errors.FromError(err)
		res := ReErrJSON{
			ReJSON: ReJSON{
				Stat:    se.Stat,
				Code:    int(se.Code),
				Message: se.Message,
				Msg:     se.Message,
				Data:    data,
				TraceID: c.GetString(TraceIDKey),
			},
			Metadata: se.Metadata,
		}
		resByte, _ := json.Marshal(res)
		c.Set("output", string(resByte))

		c.JSON(http.StatusOK, res)

		return
	}

	WriteSuccessJSON(c, data)
}

func init() {
	emptyData = map[string]interface{}{}
}
