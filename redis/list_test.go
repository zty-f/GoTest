package cache

import (
	"context"
	"fmt"
	"github.com/bytedance/sonic"
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

func UserMedalsToStrings(userMedals []UserMedalMini) ([]string, error) {
	result := make([]string, 0)
	for _, v := range userMedals {
		str, err := sonic.Marshal(v)
		if err != nil {
			return nil, err
		}
		result = append(result, string(str))
	}
	return result, nil
}

func TestStructListCache(t *testing.T) {
	userMedalMinis := []UserMedalMini{
		{Id: 1, MedalId: 1, Year: 2021},
		{Id: 2, MedalId: 2, Year: 2022},
		{Id: 3, MedalId: 3, Year: 2023},
	}
	strings, err := UserMedalsToStrings(userMedalMinis)
	if err != nil {
		return
	}
	err, count := LPush("user1", strings...)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(count)
	err, res := LGetAll("user1")
	if err != nil {
		fmt.Println(err)
		return
	}
	userMedals := make([]UserMedal, 0)
	for _, v := range res {
		userMedal := UserMedal{}
		err = sonic.Unmarshal([]byte(v), &userMedal)
		if err != nil {
			fmt.Println(err)
			return
		}
		userMedals = append(userMedals, userMedal)
	}
	fmt.Printf("%+v\n", userMedals)
}

func TestAppend(t *testing.T) {
	tmp := make([]UserMedalMini, 0)
	userMedalMinis := []UserMedalMini{
		{Id: 1, MedalId: 1, Year: 2021},
		{Id: 2, MedalId: 2, Year: 2022},
		{Id: 3, MedalId: 3, Year: 2023},
	}
	tmp = append(tmp, userMedalMinis...)
	fmt.Printf("%+v\n", tmp)
}
