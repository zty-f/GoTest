package base

import (
	"fmt"
	"testing"
)

type 自定义类型 string

type 类型别名 = string

func TestType1(t *testing.T) {
	var a 自定义类型 = "test"
	fmt.Printf("%s\n", a)
	fmt.Printf("%T\n", a) // base.自定义类型
	var b 类型别名 = "test"
	fmt.Printf("%s\n", b)
	fmt.Printf("%T\n", b) // string
	var c string = "test"
	fmt.Println(string(a) == c) // 自定义的类型和原类型不能直接比较，需要强转类型
	fmt.Println(b == c)
}
