package main

import (
	"fmt"
	"test/gorm/model"
	"testing"
	"time"
)

func TestCreate1(t *testing.T) {
	userMedal := model.UserMedal{
		StuId:      11111,
		MedalId:    23333,
		IsWear:     1,
		Year:       25,
		ExtData:    "d.ExtData",
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}
	tx := DB.Begin()
	err := tx.Table("user_medal").Create(&userMedal).Error
	tx.Commit()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", userMedal)

	mini := model.UserMedalMini{
		Id:         userMedal.Id,
		MedalId:    userMedal.MedalId,
		Year:       userMedal.Year,
		CreateTime: userMedal.CreateTime,
	}
	fmt.Printf("%+v\n", mini)
}
