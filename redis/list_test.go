package main

import (
	"context"
	"fmt"
	"testing"
)

func LPush(key string, values ...string) (error, int64) {
	var ctx = context.Background()
	// 从列表左边（头部）插入
	result, err := rd.LPush(ctx, key, values).Result()
	if err != nil {
		return err, 0
	}
	return nil, result
}

func RPush(key string, values ...string) (error, int64) {
	var ctx = context.Background()
	// 从列表左边（尾部）插入
	result, err := rd.RPush(ctx, key, values).Result()
	if err != nil {
		return err, 0
	}
	return nil, result
}

func LGetAll(key string) (error, []string) {
	var ctx = context.Background()
	// 获取列表所有元素
	result, err := rd.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		return err, nil
	}
	return nil, result
}

func TestLPush(t *testing.T) {
	// 从队列左边（头部）插入
	err, count := LPush("list1", "1", "2", "3")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(count)
}

func TestRPush(t *testing.T) {
	// 从队列右边（尾部）插入
	err, count := RPush("list2", "133344")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(count)
}

func TestLGetAll(t *testing.T) {
	// 获取队列所有元素
	err, list := LGetAll("list2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(list)
}
