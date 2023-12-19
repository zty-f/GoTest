package closure

/*
	Go函数闭包测试学习
*/

import (
	"fmt"
	"testing"
)

func newCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func Test1(t *testing.T) {
	f := newCounter()
	fmt.Println(f()) // 1
	fmt.Println(f()) // 2
	f1 := newCounter()
	fmt.Println(f1()) // 1
	fmt.Println(f1()) // 2
}
