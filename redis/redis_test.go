package cache

import (
	"codeup.aliyun.com/61e54b0e0bb300d827e1ae27/backend/golib/logger"
	"context"
	"errors"
	"fmt"
	"math/rand"
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
	v, err := rd.HGet(ctx, "yyy", "t").Result()
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

func TestHLen(t *testing.T) {
	ctx := context.Background()
	v, err := rd.HLen(ctx, "4444444").Result()
	if err != nil {
		fmt.Println(err, v)
		return
	}
	fmt.Println("-------")
	fmt.Println(v)
}

func TestExists(t *testing.T) {
	ctx := context.Background()
	v, err := rd.Exists(ctx, "zset").Result()
	if err != nil {
		fmt.Println(err, v)
		return
	}
	fmt.Println("-------")
	fmt.Println(v)
}

func TestSetNx(t *testing.T) {
	ctx := context.Background()
	v, err := rd.SetNX(ctx, "t", "c", 10*time.Second).Result()
	if err != nil {
		fmt.Println(err, v)
		return
	}
	fmt.Println("-------")
	fmt.Println(v)
}

func TestExpireAt(t *testing.T) {
	ctx := context.Background()
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(100))
	x := rand.Int63()
	v, err := rd.ExpireAt(ctx, "zset", time.Now().Add(time.Duration(x)*time.Second)).Result()
	if err != nil {
		fmt.Println(err, v)
		return
	}
	fmt.Println("-------")
	fmt.Println(v)
}

func TestExpireNX(t *testing.T) {
	ctx := context.Background()
	// redis 7.0以上版本才支持
	v, err := rd.ExpireNX(ctx, "list2", 120*time.Second).Result()
	if err != nil {
		fmt.Println(err, v)
		return
	}
	fmt.Println("-------")
	fmt.Println(v)
}

func Test取余(t *testing.T) {
	fmt.Println(2100051684 / 8 % 3)
}
