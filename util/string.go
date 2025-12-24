package util

import (
	"fmt"
	"github.com/bytedance/sonic"
	"math"
	"math/rand"
	"strings"
	"time"
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

func CutStringAfterSep(str, sep string) string {
	// 截断【限时奖励】之后的字符串
	index := strings.Index(str, sep)
	if index == -1 {
		return str
	}
	return str[:index]
}

// DiffStringArray 返回在 a 中但不在 b 中的数据
func DiffStringArray(a, b []string) []string {
	result := make([]string, 0)

	bMap := make(map[string]struct{})
	for _, v := range b {
		bMap[v] = struct{}{}
	}

	for _, v := range a {
		if _, ok := bMap[v]; !ok {
			result = append(result, v)
		}
	}

	return result
}

func ZeroIfEmpty(value string) string {
	if value == "" {
		return "0"
	}
	return value
}

// RemoveStringToString 移除逗号分隔字符串中的指定字符串
func RemoveStringToString(baseStr string, sep string) string {
	cList := strings.Split(baseStr, ",")
	if len(cList) == 0 {
		return ""
	}
	var cListNew []string
	for _, cl := range cList {
		if cl != sep {
			cListNew = append(cListNew, cl)
		}
	}
	return strings.Join(cListNew, ",")
}

// TrimPrefix0 去掉字符串前面的0
func TrimPrefix0(str string) string {
	return strings.TrimPrefix(str, "0")
}

// AddMonthPrefix0 给月份字符串前面加0
func AddMonthPrefix0(str string) string {
	if len(str) == 1 {
		return "0" + str
	}
	return str
}

// PadLeftRight 在左侧/右侧填充指定字符直到达到目标长度
func PadLeftRight(str string, padChar rune, length int, isLeft bool) string {
	strLen := len(str)
	if strLen >= length {
		return str
	}
	padding := strings.Repeat(string(padChar), length-strLen)
	if isLeft {
		return padding + str
	}
	return str + padding
}

func RandomStr(list []string) string {
	if len(list) == 0 {
		return ""
	}
	return list[rand.Intn(len(list))]
}

// GenerateRandomFloatBetween 生成指定范围内的一位小数
func GenerateRandomFloatBetween(min, max int) float64 {
	if min > max {
		min, max = max, min
	}
	rand.Seed(time.Now().UnixNano())
	randomFloat := float64(min) + rand.Float64()*(float64(max)-float64(min))
	return math.Floor(randomFloat*10) / 10
}
