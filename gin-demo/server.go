package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"math/rand"
	"time"
)

func Get(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello gin!",
	})
}

type Test struct {
	Name   string `form:"name" json:"name"`
	Brand  string `form:"brand" json:"brand"`
	UserId int    `form:"user_id" json:"user_id"`
}

type Test1 struct {
	Name         string `form:"name" json:"name"`
	Brand        string `form:"brand" json:"brand"`
	UserIdString string `form:"user_id" json:"user_id"`
}

func TestShouldBind(c *gin.Context) {
	test := &Test{}
	err := c.ShouldBindWith(test, binding.Form) //这个不会影响bind次数
	if err != nil {
		fmt.Println(err)
		return
	}
	test2 := &Test{}
	err2 := c.ShouldBind(test2)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	// ShouldBind绑定参数时，如果参数类型是form-data/x-www-form-urlencoded时,可以多次使用ShouldBind
	// 但是如果参数类型是json时，只能使用一次ShouldBind
	//test3 := &Test{}
	//err3 := c.ShouldBindJSON(test3)
	//if err3 != nil {
	//	fmt.Println(err3)
	//	return
	//}
	//test4 := &Test{}
	//err4 := c.ShouldBindJSON(test4)
	//if err4 != nil {
	//	fmt.Println(err4)
	//	return
	//}
	// ShouldBindJSON不能使用多次，尽管是针对不同地址空间的结构体,也不能和shouldBind共同使用多次
	//1.单次解析，追求性能使用 ShouldBindJson，因为多次绑定解析会出现EOF
	//2.多次解析，避免报EOF，使用ShouldBindBodyWith
	c.JSON(200, gin.H{
		"message": "hello TestShouldBind!",
	})
}

// 假设的任务函数，有一定几率失败
func task(ctx context.Context) error {
	// 模拟50%几率成功或失败
	if rand.Intn(2) == 0 {
		return fmt.Errorf("task failed")
	}
	fmt.Println(ctx.Value("x_trace_id"))
	fmt.Println("task succeeded")
	return nil
}

// 异步执行任务，带有重试逻辑
func asyncTaskWithRetry(ctx context.Context, maxRetries int, delay time.Duration, t func(ctx context.Context) error) {
	go func() {
		for retries := 0; retries < maxRetries; retries++ {
			err := t(ctx)
			if err != nil {
				fmt.Println(err)
				time.Sleep(delay) // 等待一段时间再重试
				continue
			}
			break
		}
		fmt.Println("子流程执行完成")
	}()
}

func TestRetry(c *gin.Context) {
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())
	ctx := c.Request.Context()
	ctx = context.WithValue(ctx, "x_trace_id", "123456")

	// 启动异步任务，最多重试3次，每次重试间隔1秒
	asyncTaskWithRetry(ctx, 5, 20*time.Second, task)

	// 主流程继续执行其他任务
	fmt.Println("主流程继续执行...")
	time.Sleep(2 * time.Second) // 假设主流程有其他耗时任务
	fmt.Println("主流程结束")
	c.JSON(200, gin.H{
		"code":    200,
		"message": "done",
	})
}
