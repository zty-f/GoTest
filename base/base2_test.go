package base

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
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

func TestCast(t *testing.T) {
	// cast 方法如果不是数值使用cast.ToInt会返回0
	str1 := "20019_finish"
	str2 := "2001000"
	str3 := "finishNum"
	fmt.Println(cast.ToInt(str1)) // 0
	fmt.Println(cast.ToInt(str2)) // 2001000
	fmt.Println(cast.ToInt(str3)) // 0
}

type MapData struct {
	Map  map[string]string
	Name string
	Id   int
}

func TestMapNil(t *testing.T) {
	m := &MapData{}
	m.Map["1"] = "1"
	fmt.Println(m.Map)
}

type Prize struct {
	PrizeId    string `json:"prizeId"`
	PrizeName  string `json:"prizeName"`
	PrizeType  int    `json:"prizeType"`
	PrizeCount int    `json:"prizeCount"`
	PrizeImage string `json:"prizeImage"`
}

func TestModelStr(t *testing.T) {
	prizes := make([]Prize, 0)
	prizes = append(prizes, Prize{
		PrizeId:    "86",
		PrizeName:  "蝙蝠侠",
		PrizeType:  4,
		PrizeCount: 1,
		PrizeImage: "https://static-inc.xiwang.com/mall/caa9c21cb0d27b62365d996753cacb88.png",
	})
	prizes = append(prizes, Prize{
		PrizeId:    "xxx",
		PrizeName:  "抽奖次数",
		PrizeType:  5,
		PrizeCount: 3,
		PrizeImage: "https://static-inc.xiwang.com/mall/b4f35c39ee47240a2db91f579056933d.png",
	})
	marshal, _ := json.Marshal(prizes)
	fmt.Println(string(marshal))
}
