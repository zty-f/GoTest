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

func defer_call() {

	defer func() {
		println("defer: panic 捕获之后继续按照顺序执行defer")
	}()

	defer func() {
		fmt.Println("defer: panic 之前1, 捕获异常")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	defer func() { fmt.Println("defer: panic 之前2, 不捕获") }()

	panic("异常内容") //触发defer出栈

	defer func() { fmt.Println("defer: panic 之后, 永远执行不到") }()
}

func TestDeferPanic(t *testing.T) {
	defer_call()

	fmt.Println("main 正常结束")
}
