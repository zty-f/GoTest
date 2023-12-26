package main

import (
	"fmt"
	"gorm.io/gorm/clause"
	"test/gorm/model"
	"testing"
	"time"
)

func TestUpdates(t *testing.T) {
	var usertask = model.UserTask{
		CreateTime: 11111122,
		UpdateTime: 0,
		IsDeleted:  0,
		Extra:      "",
	}
	// gorm更新不存在的记录不会报错，只是RowsAffected:0
	res := DB.Debug().Table("user_task").Where("unique_id = ?", "22222222333").Updates(&usertask)

	fmt.Printf("%+v\n", res)
	fmt.Printf("%+v\n", usertask)
}

func TestSelect(t *testing.T) {
	var usermedal model.UserMedal
	// find
	err := DB.Debug().Table("user_medal").Where("stu_id = ?", 2100051684).Find(&usermedal).Limit(1).Error
	fmt.Println(err)
	fmt.Printf("%+v\n", usermedal)

	// take
	err = DB.Debug().Table("user_medal").Where("stu_id = ?", 11111).Take(&usermedal).Error
	fmt.Println(err)
	fmt.Printf("%+v\n", usermedal)

	// Last
	err = DB.Debug().Table("user_medal").Where("stu_id = ?", 11111).Last(&usermedal).Error
	fmt.Println(err)
	fmt.Printf("%+v\n", usermedal)

	// First
	err = DB.Debug().Table("user_medal").Where("stu_id = ?", 11111).First(&usermedal).Error
	fmt.Println(err)
	fmt.Printf("%+v\n", usermedal)
}

func TestCount(t *testing.T) {
	var res []model.UserMedal
	var count int64
	// find
	err := DB.Debug().Table("user_medal").Count(&count).Limit(2).Find(&res).Error
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", res)
	fmt.Println(count)
}

type UserAvatarDecoration struct {
	Id             int   `json:"id" gorm:"column:id;primaryKey"`
	StuId          int64 `json:"stuId" gorm:"column:stu_id"`
	DecorationId   int   `json:"decorationId" gorm:"column:decoration_id"`
	DecorationType int   `json:"decorationType" gorm:"column:decoration_type"`
	CreateTime     int64 `json:"createTime" gorm:"column:create_time"`
	UpdateTime     int64 `json:"update_time" gorm:"update_time"`
	Status         int   `json:"status" gorm:"status"`
}

func TestUpsert(t *testing.T) {
	var x = &UserAvatarDecoration{
		StuId:          2100051684,
		DecorationId:   44,
		DecorationType: 1,
	}
	err := DB.Debug().Table("user_avatar_decoration").Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "stu_id"}, {Name: "decoration_type"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"decoration_id": x.DecorationId, "status": 1, "update_time": time.Now().Unix()}),
	}).Create(&x).Error
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", x)
}
