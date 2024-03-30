package util

import (
	"github.com/bytedance/sonic"
)

func InterfaceToString(i interface{}) string {
	b, err := sonic.Marshal(i)
	if err != nil {
		return ""
	}
	return string(b)
}
