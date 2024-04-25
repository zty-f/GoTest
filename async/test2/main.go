package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 假设的任务函数，有一定几率失败
func task() error {
	// 模拟50%几率成功或失败
	if rand.Intn(2) == 0 {
		return fmt.Errorf("task failed")
	}
	fmt.Println("task succeeded")
	return nil
}

// 异步执行任务，带有重试逻辑
func asyncTaskWithRetry(maxRetries int, delay time.Duration) {
	go func() {
		for retries := 0; retries < maxRetries; retries++ {
			err := task()
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

func main() {
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())

	// 启动异步任务，最多重试3次，每次重试间隔1秒
	asyncTaskWithRetry(3, 10*time.Second)

	// 主流程继续执行其他任务
	fmt.Println("主流程继续执行...")
	time.Sleep(1 * time.Second) // 假设主流程有其他耗时任务
	fmt.Println("主流程结束")
}
