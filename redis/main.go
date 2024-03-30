package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type UserMedalMini struct {
	Id         int   `gorm:"id" json:"id"`
	MedalId    int   `gorm:"medal_id" json:"medalId"`
	Year       int   `gorm:"year" json:"year"`
	CreateTime int64 `gorm:"create_time" json:"createTime"`
}

type UserMedal struct {
	Id         int    `gorm:"id" json:"id"`
	StuId      int64  `gorm:"stu_id" json:"stuId"`
	MedalId    int    `gorm:"medal_id" json:"medalId"`
	Year       int    `gorm:"year" json:"year"`
	IsWear     int    `gorm:"is_wear" json:"isWear"`
	ExtData    string `gorm:"ext_data" json:"extData"`
	WearTime   int64  `gorm:"wear_time" json:"wearTime"`
	CreateTime int64  `gorm:"create_time" json:"createTime"`
	NoticeTime int64  `gorm:"notice_time" json:"noticeTime"`
	UpdateTime int64  `gorm:"update_time" json:"updateTime"`
}

var rd *redis.Client

func main() {
	//testSetNx()
	//testSet()
	testExpireNx()
}

func init() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	rd = rdb
}

func testSetNx() {
	var ctx = context.Background()
	fmt.Println("-------------------")
	// 过期时间不能这么设置，会报错，但是插入正常，1ms过期
	result, err := rd.SetNX(ctx, "aaaaa", "11111", 86400*2).Result()
	if err != nil {
		fmt.Println(err)
	}
	println(result)
}

func testSet() {
	var ctx = context.Background()
	fmt.Println("-------------------")
	// 过期时间不能这么设置，会报错，但是插入正常，1ms过期
	result, err := rd.Set(ctx, "bbbb", "11111", 86400*2).Result()
	if err != nil {
		fmt.Println(err)
	}
	println(result)
}

func testExpireNx() {
	var ctx = context.Background()
	fmt.Println("-------------------")
	// redis 7版本才会支持expireNx，低版本不支持这个命令
	result, err := rd.Expire(ctx, "set", 100*time.Second).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	println(result)
}
