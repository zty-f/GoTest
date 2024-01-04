package util

import (
	"context"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"time"
)

// StrToUnixTime 字符串转时间戳
func StrToUnixTime(str string) (int64, error) {
	layout := "2006-01-02 15:04:05"
	t, err := time.ParseInLocation(layout, str, time.Local)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

func StrToTime(str string) (time.Time, error) {
	layout := "2006-01-02 15:04:05"
	return time.ParseInLocation(layout, str, time.Local)
}

// TimeToStr 时间戳转字符串
func TimeToStr(timer int64) string {
	tm := time.Unix(timer, 0)
	return tm.Format("2006-01-02 15:04:05")
}

// IsNil IsNil判断一个值是否为nil，特定类型已经声明但未赋值也会返回true
func IsNil(expr interface{}) bool {
	if nil == expr {
		return true
	}
	v := reflect.ValueOf(expr)
	k := v.Kind()
	return k >= reflect.Chan && k <= reflect.Slice && v.IsNil()
}

func CtxGetString(ctx context.Context, key string) (string, bool) {
	if ctx == nil || IsNil(ctx.Value(key)) {
		return "", false
	}
	if val, ok := ctx.Value(key).(*string); ok {
		return *val, ok
	}
	if val, ok := ctx.Value(key).(string); ok {
		return val, ok
	}
	return "", false
}

// GetContextUserWorkCode 获取工号
func GetContextUserWorkCode(c context.Context) string {
	val, _ := CtxGetString(c, "work_code")
	return val
}

// GetContextUserName 获取用户姓名
func GetContextUserName(c context.Context) string {
	val, _ := CtxGetString(c, "user_name")
	return val
}

// GetOperatorInfo 通用的操作用户
func GetOperatorInfo(c context.Context) string {
	name := GetContextUserName(c)
	workCode := GetContextUserWorkCode(c)
	operator := ""
	if len(workCode) == 0 {
		operator = name
	} else {
		operator = fmt.Sprintf("%s(%s)", name, workCode)
	}
	return operator
}

// GetNumberWithBitsSet 将指定位数变为1
func getNumberWithBitsSet1(nums []int) int {
	num := 0
	for _, i := range nums {
		if i <= 0 {
			return 0
		}
		x := math.Pow(2, float64(i-1))
		num = num + int(math.Round(x))
	}
	return num
}

func getNumberWithBitsSet2(positions ...int) int {
	var result int
	for _, position := range positions {
		result |= 1 << (position - 1)
	}
	return result
}

// 取一个十进制数对应的二进制数哪些位等于1
func getBinaryBits1(num int) ([]int, error) {
	number := strconv.FormatInt(int64(num), 2)
	res, err := strconv.ParseUint(number, 2, 64)
	if err != nil {
		return nil, err
	}
	i := 0
	result := []int{}
	for res != 0 {
		i++
		if res&1 == 1 {
			result = append(result, i)
		}
		res = res >> 1
	}
	return result, nil
}

func getBinaryBits2(num int) []int {
	var bits []int
	position := 1
	for num > 0 {
		if num&1 == 1 {
			bits = append(bits, position)
		}
		num >>= 1
		position++
	}
	return bits
}

// GetMonthStartTime 获取某个月的开始时间戳 dateStr：2023-10
func GetMonthStartTime(dateStr string) (int64, error) {
	layout := "2006-01"
	date, err := time.Parse(layout, dateStr)
	if err != nil {
		return 0, err
	}
	year, month, _ := date.Date()
	loc, _ := time.LoadLocation("Local")
	startOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, loc)
	return startOfMonth.Unix(), nil
}

// GetMonthEndTime 获取某个月的结束时间戳 dateStr：2023-10
func GetMonthEndTime(dateStr string) (int64, error) {
	layout := "2006-01"
	date, err := time.Parse(layout, dateStr)
	if err != nil {
		return 0, err
	}
	year, month, _ := date.Date()
	loc, _ := time.LoadLocation("Local")
	nextMonth := time.Date(year, month+1, 1, 0, 0, 0, 0, loc)
	endOfMonth := nextMonth.Add(-time.Second)
	return endOfMonth.Unix(), nil
}
