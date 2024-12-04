package util

import (
	"math/rand"
	"time"
)

func GetRandomOneFromStringArray(arr []string) string {
	if len(arr) == 0 {
		return ""
	}
	// 初始化随机种子
	rand.Seed(time.Now().UnixNano())
	return arr[rand.Intn(len(arr))]
}
