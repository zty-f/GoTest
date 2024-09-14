package base

import (
	"fmt"
	"testing"
	"time"
)

// 两个时间范围的交集判断
func Test_1(t *testing.T) {
	time1 := time.Now()
	time2 := time.Now().Add(7 * time.Hour * 24)
	beginTime := time.Unix(1726279602, 0)
	endTime := time.Unix(1726379602, 0)
	// 如果time1到time2的时间范围内存在一个时间点在beginTime和endTime之间，就返回true，否则返回false
	// 上述情况归根结底就是两个范围判断交集的问题
	if hasOverlap(time1, time2, beginTime, endTime) {
		fmt.Println("The time ranges overlap")
	} else {
		fmt.Println("The time ranges do not overlap")
	}
}

func hasOverlap(time1, time2, beginTime, endTime time.Time) bool {
	return (time1.Before(endTime) || time1.Equal(endTime)) && (time2.After(beginTime) || time2.Equal(beginTime))
}
