package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// 一个模拟的操作，有时会失败
func doSomething(ctx context.Context) error {
	// 随机失败
	if rand.Intn(10) < 5 {
		return fmt.Errorf("operation failed")
	}
	return nil
}

// 异步重试函数
func asyncRetry(ctx context.Context, maxRetries int, retryDelay time.Duration, operation func(ctx context.Context) error) <-chan error {
	// 创建一个channel用来返回结果
	resultChan := make(chan error, 1)

	go func() {
		defer close(resultChan)
		for i := 0; i < maxRetries; i++ {
			// 尝试执行操作
			err := operation(ctx)
			if err == nil {
				// 成功，返回结果
				resultChan <- nil
				return
			}
			// 判断context是否已经结束
			select {
			case <-ctx.Done():
				// context结束，返回错误
				resultChan <- ctx.Err()
				return
			case <-time.After(retryDelay):
				// 等待一段时间后重试
			}
		}
		// 重试次数达到最大值后仍然失败，返回最后的错误
		resultChan <- fmt.Errorf("operation failed after retries")
	}()

	return resultChan
}

func main() {
	// 初始化随机数生成器
	rand.Seed(time.Now().UnixNano())

	// 设置最大重试次数和重试间隔
	maxRetries := 3
	retryDelay := 2 * time.Second

	// 创建一个超时时间为5秒的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 调用异步重试函数
	resultChan := asyncRetry(ctx, maxRetries, retryDelay, doSomething)

	// 等待结果
	err := <-resultChan
	if err != nil {
		fmt.Println("Operation failed:", err)
	} else {
		fmt.Println("Operation succeeded.")
	}
}
