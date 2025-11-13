package limit

import (
	"context"
	"fmt"
	"golang.org/x/time/rate"
	"log"
	"testing"
	"time"
)

// https://juejin.cn/post/7258233838370603069
/*
使用NewLimiter创建一个限流器Limiter：
r：表示速率，每秒产生r个令牌
b：表示桶大小，最大突发b个事件
示例代码：
如下表示限制10 QPS，突发1
limiter := NewLimiter(10, 1);
*/
// 官方令牌桶限流器使用
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

func TestRate2(t *testing.T) {
	r := rate.NewLimiter(1, 5)
	for {
		allow := r.AllowN(time.Now(), 2)
		if allow {
			fmt.Println(time.Now().Format("2021-11-02 15:04:05") + " allow")
		} else {
			fmt.Println(time.Now().Format("2021-11-02 15:04:05") + " not allow")
		}
		time.Sleep(time.Second)
	}
}

func TestTokenBucket_Take(t *testing.T) {
	tb := NewTokenBucket(5, 1)
	for {
		allow := tb.Take()
		if allow {
			fmt.Println(time.Now().Format("2021-11-02 15:04:05") + " allow")
		} else {
			fmt.Println(time.Now().Format("2021-11-02 15:04:05") + " not allow")
		}
		time.Sleep(time.Second)
	}
}

func TestTokenBucket_TakeN(t *testing.T) {
	tb := NewTokenBucket(5, 1)
	for {
		allow := tb.TakeN(2)
		if allow {
			fmt.Println(time.Now().Format("2021-11-02 15:04:05") + " allow")
		} else {
			fmt.Println(time.Now().Format("2021-11-02 15:04:05") + " not allow")
		}
		time.Sleep(time.Second)
	}
}
