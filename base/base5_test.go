package base

import (
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
