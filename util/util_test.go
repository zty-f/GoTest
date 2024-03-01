package util

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func Test1(t *testing.T) {
	//num := 0007
	//number := strconv.FormatInt(int64(num), 2)
	//fmt.Println(number)
	number := "42"
	res, err := strconv.ParseUint(number, 2, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
	// ParseInt实质上先对于字符串数字前面的+-符号进行处理，然后再调用ParseUint方法
	res1, err := strconv.ParseInt(number, 2, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res1)
}

func TestGetMonthStartTime(t *testing.T) {
	dateStr := "2024-01"
	time, _ := GetMonthStartTime(dateStr)
	fmt.Println(time)
	time, _ = GetMonthEndTime(dateStr)
	fmt.Println(time)
}

// 生成随机字符串
var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}
func GetRandomString1(l int) string {
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	var result []byte
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func GetRandomString2(l int) string {
	bytes := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	var result = make([]byte, l)
	for i := 0; i < l; i++ {
		result[i] = bytes[r.Intn(len(bytes))]
	}
	return string(result)
}

func GetRandomString3(n int) string {
	randBytes := make([]byte, n/2)
	rand.Read(randBytes)
	return fmt.Sprintf("%x", randBytes)
}

func TestGetRandomString(t *testing.T) {
	fmt.Println(GetRandomString1(10))
	fmt.Println(GetRandomString2(10))
	fmt.Println(GetRandomString3(10))
}

func BenchmarkGetRandomString1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetRandomString1(10)
	}
}

func BenchmarkGetRandomString2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetRandomString2(10)
	}
}

func BenchmarkGetRandomString3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetRandomString3(10)
	}
}
