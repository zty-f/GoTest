package main

import (
	"fmt"
	"test/gorm/model"
	"testing"
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
