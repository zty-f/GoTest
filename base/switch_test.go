package base

import (
	"regexp"
	"strconv"
	"strings"
	"testing"
)

// 解码八进制转义字符串为UTF-8中文
func OctalEscapeToUTF8(octalStr string) (string, error) {
	// 方法A: 使用正则匹配替换
	re := regexp.MustCompile(`\\(\d{3})`)
	result := re.ReplaceAllStringFunc(octalStr, func(match string) string {
		octalNum := match[1:] // 去掉开头的 \
		val, _ := strconv.ParseUint(octalNum, 8, 8)
		return string([]byte{byte(val)})
	})

	// 结果已经是UTF-8字节序列
	return result, nil
}

// 支持多种转义格式
func DecodeEscapedString(escapedStr string) (string, error) {
	// 处理八进制 \346
	if strings.Contains(escapedStr, "\\") {
		return OctalEscapeToUTF8(escapedStr)
	}

	// 处理十六进制 \xE6
	if strings.Contains(escapedStr, "\\x") {
		quoted := `"` + escapedStr + `"`
		return strconv.Unquote(quoted)
	}

	return escapedStr, nil
}

func TestOctalEscapeToUTF8(t *testing.T) {
	octalStr := "\\344\\271\\260\\344\\270\\215\\344\\271\\260" // "测试" 的八进制转义
	utf8Str, err := OctalEscapeToUTF8(octalStr)
	if err != nil {
		panic(err)
	}
	println("Octal to UTF-8:", utf8Str)
}
