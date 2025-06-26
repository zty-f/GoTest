package base

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"testing"
	"time"
)

func TestNilStruct(t *testing.T) {
	type User struct {
		Name string
	}

	var u *User
	if u == nil {
		t.Log("u is nil")
	} else {
		t.Error("u should be nil")
	}

	u = &User{}

	fmt.Println(u.Name)
	fmt.Println(u.Name)
	fmt.Println(u.Name)
	fmt.Println(cast.ToString(0))

}

func TestThreeMonthAgo(t *testing.T) {
	// 获取当前时间
	now := time.Now()

	// 计算三个月前的时间
	threeMonthsAgo := now.AddDate(0, -3, 0)

	// 打印结果
	fmt.Println("当前时间:", now)
	fmt.Println("三个月前的时间:", threeMonthsAgo)
}

type UserReportConfig struct {
	Reasons []string `json:"reasons"`
}

func TestMarshal(t *testing.T) {
	config := UserReportConfig{
		Reasons: []string{"1", "2", "3"},
	}
	marshal, _ := json.Marshal(config)
	fmt.Printf("%+v\n", string(marshal))
	// {
	//			ReportType: 1,
	//			Reasons:    []string{"垃圾营销", "人身攻击", "淫秽色情", "发布违规内容", "发布其他不适当内容"},
	//		},
	configs := map[int]UserReportConfig{
		1: {
			Reasons: []string{"垃圾营销", "人身攻击", "淫秽色情", "发布违规内容", "发布其他不适当内容"},
		},
	}
	marshal, _ = json.Marshal(configs)
	fmt.Printf("%+v\n", string(marshal))
}

type PrivacyType int64

type PrivacyConfig struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func TestPrivacyTypeMarshal(t *testing.T) {
	config := map[PrivacyType][]*PrivacyConfig{
		1: {
			{Id: 1, Name: "学习时长"},
			{Id: 2, Name: "关注列表"},
			{Id: 3, Name: "关注者列表"},
			{Id: 4, Name: "毕业证书"},
		},
	}
	marshal, _ := json.Marshal(config)
	fmt.Printf("%+v\n", string(marshal))
}
