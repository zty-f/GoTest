package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	layout := "2006-01-02 15:04:05"
	now, _ := time.ParseInLocation(layout, "2023-11-20 15:04:05", time.Local)
	fmt.Println(now.Unix())

	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(10)
	fmt.Println("随机值：", x)

	now = now.Add(-time.Duration(x) * time.Second)
	fmt.Println(now.Unix())
}
