package main

import (
	"fmt"
	"time"
)

// 根据过期剩余时间秒数获取具体截止日期
func main() {
	ttl := int64(1840081)
	dateTime := time.Now().Unix() + ttl

	fmt.Println("时间戳：", dateTime)

	tm := time.Unix(dateTime, 0)
	date := tm.Format("2006-01-02 15:04:05")
	fmt.Println("日期：", date)
}
