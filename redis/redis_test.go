package main

import (
	"codeup.aliyun.com/61e54b0e0bb300d827e1ae27/backend/golib/logger"
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

func TestRedisContextDeadlineExceeded(t *testing.T) {
	// 创建一个Redis客户端
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis服务器的地址和端口
		Password: "",               // Redis服务器的密码，如果没有设置密码则为空字符串
		DB:       0,                // Redis数据库的索引
	})

	// 创建一个上下文对象并设置超时时间为1秒   指定时间内没有完成操作就会报错context deadline exceeded
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	// 执行GET操作，并传入上下文对象
	value, err := client.Get(ctx, "b").Result()
	if err == redis.Nil {
		fmt.Println("key does not exist")
	} else if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("value:", value)
	}

	// 处理超时错误
	if errors.Is(err, context.DeadlineExceeded) {
		fmt.Println("Redis context deadline exceeded")
	} else if err != nil {
		fmt.Println("error:", err)
	}
}

func TestHGet(t *testing.T) {
	ctx := context.Background()
	v, err := rd.HGet(ctx, "a", "b").Result()
	if err != nil {
		if err != redis.Nil {
			logger.Ex(ctx, "redisHGetFailed", "%s", err.Error())
		}
		fmt.Println(err, v)
		return
	}
	fmt.Println("-------")
	fmt.Println(v)
}

func TestHSetNx(t *testing.T) {
	ctx := context.Background()
	v, err := rd.HSetNX(ctx, "a", "c", "2").Result()
	if err != nil {
		fmt.Println(err, v)
		return
	}
	fmt.Println("-------")
	fmt.Println(v)
}
