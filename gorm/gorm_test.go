package main

import (
	"fmt"
	"gorm.io/gorm/clause"
	"strings"
	"test/gorm/model"
	"testing"
	"time"
)

func TestUpdates(t *testing.T) {
	var usertask = model.UserTask{
		CreateTime: 11111122,
		UpdateTime: 0,
		IsDeleted:  0,
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

// 1.通过ON DUPLICATE KEY来实现创建或更新 需要有唯一主键或者唯一索引才能支持，一条sql
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

// 2.通过FirstOrCreate 来实现创建或更新 先查询 然后再根据情况创建或更新，两条sql
func TestUpsert3(t *testing.T) {
	var x = &UserAvatarDecoration{
		StuId:          2100051684,
		DecorationId:   44,
		DecorationType: 1,
	}
	var create = x
	err := DB.Table("user_avatar_decoration").Where("stu_id = ? and type = ?", 111, 1).Assign(&create).FirstOrCreate(&x).Error
	if err != nil {
		fmt.Println(err)
	}
}

// ErrorsIsDuplicate 判断错误是否是数据重复错误
func ErrorsIsDuplicate(err error) bool {
	if err == nil {
		return false
	}

	if strings.Index(err.Error(), "Duplicate") != -1 {
		return true
	}

	return false
}

func TestDuplicateErr(t *testing.T) {
	var x = &model.UserTask{
		TaskId:     2001000,
		StuId:      2100051764,
		CreateTime: 0,
		UpdateTime: 0,
		UniqueId:   "2100051764-2001000-1-7",
		IsDeleted:  0,
	}
	err := DB.Table("user_task_0").Create(&x).Error
	fmt.Println(err)
	if ErrorsIsDuplicate(err) {
		fmt.Println("true")
	}
}

type User struct {
	Id   int
	Name string
}

func TestSelectIn1(t *testing.T) {
	ids1 := []int{2, 3}
	users1 := make([]User, 0)
	err := DB.Table("user_1").Find(&users1, ids1).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("users:%+v\n", users1)

	var ids2 []int
	users2 := make([]User, 0)
	err = DB.Table("user_1").Find(&users2, ids2).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("users:%+v\n", users2)
}

func TestSelectIn2(t *testing.T) {
	ids1 := []int{2, 3}
	users1 := make([]User, 0)
	err := DB.Table("user_1").Find(&users1, "id in (?)", ids1).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("users:%+v\n", users1)

	var ids2 []int
	users2 := make([]User, 0)
	err = DB.Table("user_1").Find(&users2, "id in (?)", ids2).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("users:%+v\n", users2)
}

func TestSort(t *testing.T) {
	users1 := make([]User, 0)
	// 如下语句的排序字段等同于name asc,id desc，不写值会默认升序
	err := DB.Debug().Table("user_1").Order("name,id desc").Find(&users1).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("users1:%+v\n", users1)
	users2 := make([]User, 0)
	err = DB.Debug().Table("user_1").Order("name desc,id desc").Find(&users2).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("users2:%+v\n", users2)
}

func TestCreate2(t *testing.T) {
	userss := make([]User, 0)
	// user1 := User{
	//	Name: "test",
	// }
	// user2 := User{
	//	Name: "test2",
	// }
	// userss = append(userss, user1, user2)
	// 如下语句的排序字段等同于name asc,id desc，不写值会默认升序
	err := DB.Debug().Table("user").Create(&userss).Error
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	fmt.Printf("users:%#v\n", userss)
}

func TestFind(t *testing.T) {
	var userks []*model.UserTask
	// 如下语句的排序字段等同于name asc,id desc，不写值会默认升序
	err := DB.Debug().Table("user").Where("name", "eerr").Find(&userks).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := range userks {
		fmt.Printf("userks:%+v\n", userks[i])
	}
	fmt.Printf("userks:%+v\n", userks)
}

func TestSelectWhere(t *testing.T) {
	userMedels := make([]model.UserMedal, 0)
	err := DB.Debug().Table("user_medal").Where("year = ?", 24).Where("is_wear = ? or notice_time = ?", 1, 0).Find(&userMedels).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	/*
		SQL: SELECT * FROM `user_medal` WHERE `year` = 24 AND (`is_wear` = 1 OR `notice_time` = 0)
	*/
	fmt.Printf("userMedels:%+v\n", userMedels)
}
