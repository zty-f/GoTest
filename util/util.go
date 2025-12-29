package util

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"regexp"
	"runtime/debug"
	"strconv"
)

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

// 转换类型为string
func Strval(value interface{}) string {
	var res string
	if value == nil {
		return res
	}

	switch temp := value.(type) {
	case float64:
		res = strconv.FormatFloat(temp, 'f', -1, 64)
	case float32:
		res = strconv.FormatFloat(float64(temp), 'f', -1, 64)
	case int:
		res = strconv.Itoa(temp)
	case uint:
		res = strconv.Itoa(int(temp))
	case int8:
		res = strconv.Itoa(int(temp))
	case uint8:
		res = strconv.Itoa(int(temp))
	case int16:
		res = strconv.Itoa(int(temp))
	case uint16:
		res = strconv.Itoa(int(temp))
	case int32:
		res = strconv.Itoa(int(temp))
	case uint32:
		res = strconv.Itoa(int(temp))
	case int64:
		res = strconv.FormatInt(temp, 10)
	case uint64:
		res = strconv.FormatUint(temp, 10)
	case string:
		res = temp
	case []byte:
		res = string(temp)
	default:
		newValue, _ := json.Marshal(value)
		res = string(newValue)
	}

	return res
}

// cbc解密
func CbcDecode(data, key, iv string) (string, error) {
	_data, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	_key := []byte(key)
	_iv := []byte(iv)

	block, err := aes.NewCipher(_key)
	if err != nil {
		return "", err
	}
	mode := cipher.NewCBCDecrypter(block, _iv)
	mode.CryptBlocks(_data, _data)
	_data = PKCS7UnPadding(_data)

	return string(_data), nil
}

func PKCS7Padding(data []byte) []byte {
	padding := aes.BlockSize - len(data)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

func PKCS7UnPadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}

func CheckStringForUnmarshal(checkStr string) (err error) {
	return nil
}

var EmptyErrMsg = errors.New("empty resp")

func RemoveOne(baseSlice []string, target string) []string {
	var diff []string
	for _, item := range baseSlice {
		if item != target {
			diff = append(diff, item)
		}
	}
	return diff
}

// Contains 判断某个元素是否在数组中 泛型实现
func Contains[T comparable](arr []T, item T) bool {
	for _, v := range arr {
		if v == item {
			return true
		}
	}
	return false
}

// NumToChinese 数字转中文数字
func NumToChinese(num int) (string, error) {
	if num < 0 || num > 10 {
		return "", fmt.Errorf("number out of range, should be between 0 and 10")
	}
	chineseNumbers := []string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九", "十"}
	return chineseNumbers[num], nil
}

// ChineseToNum 中文数字转数字
func ChineseToNum(chinese string) (int, error) {
	chineseNumbers := map[string]int{
		"零": 0,
		"一": 1,
		"二": 2,
		"三": 3,
		"四": 4,
		"五": 5,
		"六": 6,
		"七": 7,
		"八": 8,
		"九": 9,
		"十": 10,
	}
	if num, ok := chineseNumbers[chinese]; ok {
		return num, nil
	}
	return 0, fmt.Errorf("invalid chinese number: %s", chinese)
}

// ExtractChineseNumber 从描述中提取“第几周”里的中文数字
func ExtractChineseNumber(description string) (string, error) {
	// 匹配“第”和“周”之间的中文数字
	re := regexp.MustCompile(`第([零一二三四五六七八九十]+)周`)
	matches := re.FindStringSubmatch(description)
	if len(matches) < 2 {
		return "", fmt.Errorf("no chinese number found in 'week' description")
	}
	return matches[1], nil
}

// ExtractAndIncrementChineseWeek 从描述中提取“第几周”里的中文数字并加一替换
func ExtractAndIncrementChineseWeek(description string) (string, error) {
	// 匹配“第”和“周”之间的中文数字
	re := regexp.MustCompile(`第([零一二三四五六七八九十]+)周`)
	matches := re.FindStringSubmatch(description)
	if len(matches) < 2 {
		return "", fmt.Errorf("no chinese number found in 'week' description")
	}
	chineseNum := matches[1]
	// 中文数字转整数
	num, err := ChineseToNum(chineseNum)
	if err != nil {
		return "", err
	}
	// 整数加一
	num++
	// 检查是否超出范围
	if num > 10 {
		return "", fmt.Errorf("number out of range after increment, should be between 0 and 10")
	}
	// 整数转回中文数字
	newChineseNum, err := NumToChinese(num)
	if err != nil {
		return "", err
	}
	// 替换原描述中的中文数字
	return re.ReplaceAllString(description, "第"+newChineseNum+"周"), nil
}

// 数组转 any数组
func SliceToAny[T int | int32 | int64 | string](s []T) []any {
	anyList := make([]any, 0, len(s))
	for _, v := range s {
		anyList = append(anyList, v)
	}
	return anyList
}

// 数组转字符串数组
func SliceToString[T int | int32 | int64](s []T) []string {
	strList := make([]string, 0, len(s))
	for _, v := range s {
		strList = append(strList, strconv.FormatInt(int64(v), 10))
	}
	return strList
}

// 数组转 int64数组
func SliceToInt64[T int | int32](s []T) []int64 {
	anyList := make([]int64, 0, len(s))
	for _, v := range s {
		anyList = append(anyList, int64(v))
	}
	return anyList
}

func SliceToAnyInt[T, S int | int32 | int64](s []T) []S {
	anyList := make([]S, 0, len(s))
	for _, v := range s {
		anyList = append(anyList, S(v))
	}
	return anyList
}

func SafeGo(ctx context.Context, fun func()) {
	defer func(ctx context.Context) {
		if err := recover(); err != nil {
			stack := string(debug.Stack())
			fmt.Println(stack)
			// xlog.Errorf(ctx, "SafeGo,Goroutine Recover: %+v, stack is %s", err, stack)
		}
	}(ctx)
	fun()
}

func RandomArray[T any](list []T) T {
	if len(list) == 0 {
		return *new(T)
	}
	return list[rand.Intn(len(list))]
}
