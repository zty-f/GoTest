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

func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}

func DeferFunc4() (t int) {
	defer func(i int) {
		fmt.Println(i)
		fmt.Println(t)
	}(t)
	t = 1
	return 2
}

func TestDeferSum(t *testing.T) {
	fmt.Println(DeferFunc1(1))
	fmt.Println(DeferFunc2(1))
	fmt.Println(DeferFunc3(1))
	DeferFunc4()
}
