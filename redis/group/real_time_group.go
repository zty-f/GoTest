package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/cast"
	"log"
	"strconv"
	"sync"
)

var ctx = context.Background()

// 将用户添加到组中
func addUserToGroup(rdb *redis.Client, userID string) error {
	// 获取分布式锁
	lockKey := "addUserLock"
	lock, err := rdb.SetNX(ctx, lockKey, "locked", 0).Result()
	if err != nil {
		return err
	}
	if !lock {
		return errors.New("failed to acquire lock")
	}
	defer rdb.Del(ctx, lockKey) // 释放锁

	result, err := rdb.Get(ctx, "current_group_index").Result()
	if err != nil && err != redis.Nil {
		return err
	}
	curGroupIndex := 1
	if result != "" {
		curGroupIndex, err = strconv.Atoi(result)
		if err != nil {
			return err
		}
	}
	groupKey := fmt.Sprintf("group_%d", curGroupIndex)
	n, err := rdb.SCard(ctx, groupKey).Result()
	if err != nil && err != redis.Nil {
		return err
	}
	if n >= 30 {
		str, err := rdb.Incr(ctx, "current_group_index").Result()
		if err != nil {
			return err
		}
		curGroupIndex = cast.ToInt(str)
		groupKey = fmt.Sprintf("group_%d", curGroupIndex)
	}
	rdb.SAdd(ctx, groupKey, userID)
	if err != nil {
		return err
	}
	return nil
}

// 获取组中的所有用户
func getUsersInGroup(rdb *redis.Client, groupKey string) ([]string, error) {
	users, err := rdb.SMembers(ctx, groupKey).Result()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(x int) {
			for j := 1; j <= 100; j++ {
				for {
					userID := fmt.Sprintf("user_%d_%d", x, j)
					err := addUserToGroup(rdb, userID)
					if err == nil {
						break
					}
				}
			}
			wg.Done()
		}(i)
	}

	wg.Wait()

	// 打印每个组中的用户
	currentIndex, err := rdb.Get(ctx, "current_group_index").Int()
	if err != nil {
		log.Fatalf("无法获取当前组索引: %v", err)
	}

	for i := 1; i <= currentIndex; i++ {
		groupKey := fmt.Sprintf("group_%d", i)
		users, err := getUsersInGroup(rdb, groupKey)
		if err != nil {
			log.Fatalf("无法获取组 %s 中的用户: %v", groupKey, err)
		}
		fmt.Printf("组 %d: %v\n", i, users)
	}
}
