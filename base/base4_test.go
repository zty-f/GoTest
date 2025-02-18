package base

import (
	"fmt"
	"net/http"
	"testing"
	"time"
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

	panic("异常内容") // 触发defer出栈

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

func TestTimer(t *testing.T) {
	ch1 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 2)
		ch1 <- "hello"
	}()

	select {
	case res := <-ch1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout")
	}
}

func TestTimer2(t *testing.T) {
	fmt.Println("start...")
	ch1 := make(chan string, 120)

	go func() {
		time.Sleep(time.Second * 10)
		i := 0
		for {
			i++
			ch1 <- fmt.Sprintf("%s %d", "hello", i)
		}

	}()

	go func() {
		// http 监听8080, 开启 pprof
		if err := http.ListenAndServe(":8080", nil); err != nil {
			fmt.Println("listen failed")
		}
	}()
	// time.After 放到 for 外面
	timeout := time.After(time.Second * 3) // 只能使用一次
	for {
		select {
		case res := <-ch1:
			fmt.Println(res)
		case <-timeout:
			fmt.Println("timeout")
			// return
		}
	}
}
