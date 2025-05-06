package base

import (
	"context"
	"fmt"
	"golang.org/x/time/rate"
	"log"
	"net/http"
	"net/url"
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

func TestContextTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	go func() {
		select {
		case <-ctx.Done(): // 阻塞等待
			fmt.Println("timeout之后会调用这里")
			fmt.Println(ctx.Err())
		}
		fmt.Println(ctx.Err()) // context deadline exceeded
	}()

	time.Sleep(time.Second * 5)
}

func TestContextDeedLine(t *testing.T) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*1))
	defer cancel()

	go func() {
		select {
		case <-ctx.Done(): // 阻塞等待
			fmt.Println("Deadline之后会调用这里")
			fmt.Println(ctx.Err())
		}
		fmt.Println(ctx.Err()) // context deadline exceeded
	}()

	time.Sleep(time.Second * 5)
}

func TestMapInit(t *testing.T) {
	m := make(map[int][]int)
	fmt.Println(m[1])
	m[1] = append(m[1], 1)
	m[1] = append(m[1], 2)
	fmt.Println(m)
}

func TestSliceInit(t *testing.T) {
	m := make([][]int, 0)
	fmt.Println(m)
	m = append(m, make([]int, 0))
	fmt.Println(m[0])
	fmt.Println(len(m))
	// 限流
	limiter := rate.NewLimiter(1, 1)
	limiter.Wait(context.Background())
}

// https://juejin.cn/post/7258233838370603069
/*
使用NewLimiter创建一个限流器Limiter：
r：表示速率，每秒产生r个令牌
b：表示桶大小，最大突发b个事件
示例代码：
如下表示限制10 QPS，突发1
limiter := NewLimiter(10, 1);
*/
func TestRate(t *testing.T) {
	// 1表示每次放进筒内的数量，桶内的令牌数是5，最大令牌数也是5，这个筒子是自动补充的，你只要取了令牌不管你取多少个，这里都会在每次取完后自动加1个进来，因为我们设置的是1
	r := rate.NewLimiter(1, 5)
	ctx := context.Background()

	for {
		// 每次消耗2个，放入一个，消耗完了还会放进去，因为初始是5个，所以这段代码再执行到第4次的时候筒里面就空了，如果当前不够取两个了，本次就不取，再放一个进去，然后返回false
		err := r.WaitN(ctx, 2)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(time.Now().Format("2021-11-02 15:04:05"))
		time.Sleep(time.Second)
	}

}

/*
AllowN方法
AllowN 方法表示，截止到某一时刻，目前桶中数目是否至少为 n 个，满足则返回 true，同时从桶中消费 n 个 token。反之不消费桶中的Token，返回false。
此外还有Allow方法，含义和作用等同于Allow(time.Now(),1)
*/
func Limiter(next http.Handler) http.Handler {
	r := rate.NewLimiter(1, 5)
	// 这里使用http.HandlerFunc进行类型转换把匿名函数转换成了type http.HandlerFunc，因为HandlerFunc实现了ServeHttp方法所以是http.Handler的实例
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if !r.Allow() {
			http.Error(writer, "too many requests", http.StatusTooManyRequests)
		} else {
			// 获取令牌之后，再调用next.ServerHTTP继续完成请求
			next.ServeHTTP(writer, request)
		}
	})
}

type Y struct {
	Name string
	Age  int
}

func TestStructNil(t *testing.T) {
	var y *Y // nil指针
	if y == nil {
		fmt.Println("y is nil")
	}
	fmt.Println(y)
}

func TestIfElse(t *testing.T) {
	x, y := 2, 4
	// if else 语句只会执行最先满足条件的分支，后面的分支即时符合条件也不会执行
	if x > 0 {
		fmt.Println("x>0")
	} else if y > 0 {
		fmt.Println("y>0")
	} else {
		fmt.Println("x<0")
		fmt.Println("y<0")
	}
}

func TestSwitch3(t *testing.T) {
	x := 1
	switch x {
	case 1:
		fmt.Println("x=1")
	case 2:
		fmt.Println("x=2")
	default:
		fmt.Println("x!=1,2")
	}
	fmt.Println("------------")
}

func TestParseQuery(t *testing.T) {
	rawUrl := "/base/readcamp/market/track/vivo?imei=UNKNOWN&oaid=b98a037c6dac84fe2574e93d2b27c49e&timestamp=1746532109868&adid=36540213&creative_id=151603710&request_id=4fdd56e1d0484808a60eab645147aec1&ad_name=%E4%BC%B4%E9%B1%BC%E9%98%85%E8%AF%BB%E8%90%A5-%E6%B3%A8%E5%86%8C-%E6%8E%A8%E8%8D%90-%E5%AD%A9%E5%AD%90%E5%B9%B4%E9%BE%840-6%E5%B2%81-0403"
	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		fmt.Println(err)
		return
	}
	queryParam, err := url.ParseQuery(parsedUrl.RawQuery)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", queryParam)
}
