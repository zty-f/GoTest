package util

import (
	"fmt"
	"github.com/bytedance/sonic"
)

func InterfaceToString(i interface{}) string {
	b, err := sonic.Marshal(i)
	if err != nil {
		return ""
	}
	return string(b)
}

// PadLeftZeros 用零填充到指定长度
func PadLeftZeros(s string, length int) string {
	return fmt.Sprintf("%0*s", length, s)
}
