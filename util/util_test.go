package util

import (
	"fmt"
	"strconv"
	"testing"
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
