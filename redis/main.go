package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var rd *redis.Client
var ctx = context.Background()

func main() {
	testSetNx()
	testSet()
}

func init() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	rd = rdb
}

func testSetNx() {
	fmt.Println("-------------------")
	// 过期时间不能这么设置，会报错，但是插入正常，1ms过期
	result, err := rd.SetNX(ctx, "aaaaa", "11111", 86400*2).Result()
	if err != nil {
		fmt.Println(err)
	}
	println(result)
}

func testSet() {
	fmt.Println("-------------------")
	// 过期时间不能这么设置，会报错，但是插入正常，1ms过期
	result, err := rd.Set(ctx, "bbbb", "11111", 86400*2).Result()
	if err != nil {
		fmt.Println(err)
	}
	println(result)
}
