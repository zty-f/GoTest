package main

import (
	"fmt"
	"testing"
	"time"
)

var users = []User{
	{
		Name: "123",
	},
	{
		Name: "456",
	},
	{
		Name: "789",
	},
}

func TestMultiInsert(t *testing.T) {
	go Insert1()
	go Insert2()
	time.Sleep(5 * time.Second)
	/*
			结论：当两个事务同时进行插入数据的时候，数据表ID采用自增的情况下，假如有事务回滚的情况下，ID会出现不连续的情况
		    A事务 生成  1 2 3 回滚
			B事务 生成  4 5 6 正常提交
			此时数据库中的最终状态是4 5 6的情况，后续继续插入数据ID从7开始自增
			insert的时候会通过排它锁来保证ID的连续性。
	*/
}

func Insert1() {
	tx := DB.Begin()
	tx = tx.Debug()
	fmt.Println("insert1 start")
	var count int64
	tx.Table("user").Select("*").Count(&count)
	fmt.Println("insert1前", count)
	for _, user := range users {
		if err := tx.Table("user").Create(&user).Error; err != nil {
			tx.Rollback()
			return
		}
	}
	time.Sleep(3 * time.Second)
	tx.Rollback()
	var count1 int64
	DB.Table("user").Select("*").Count(&count1)
	fmt.Println("insert1后", count1)
	fmt.Println("insert1 end")
	return
	tx.Commit()
}

func Insert2() {
	tx := DB.Begin()
	tx = tx.Debug()
	fmt.Println("insert2 start")
	var count int64
	tx.Table("user").Select("*").Count(&count)
	fmt.Println("insert2前", count)
	for _, user := range users {
		if err := tx.Table("user").Create(&user).Error; err != nil {
			tx.Rollback()
			return
		}
	}
	var count1 int64
	tx.Table("user").Select("*").Count(&count1)
	fmt.Println("insert2后", count1)
	time.Sleep(3 * time.Second)
	tx.Commit()
	fmt.Println("insert2 end")
}
