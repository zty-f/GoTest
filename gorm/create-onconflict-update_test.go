package main

import (
	"fmt"
	"gorm.io/gorm/clause"
	"test/gorm/model"
	"testing"
	"time"
)

// 1.通过ON DUPLICATE KEY来实现创建或更新 需要有唯一主键或者唯一索引才能支持，一条sql
// 会导致主键不连续，每次执行的时候主键都会递增，不管本次操作是插入还是更新
/*
eg: INSERT INTO `user_monthly_report` (`stu_id`,`year_month`,`version`,`read_status`,`data`,`create_time`,`update_time`)
	VALUES (2100051684,'202401','v0.0.1',1,'{"key":"va1111lue"}',1724744553,1724744553)
	ON DUPLICATE KEY UPDATE `data`=VALUES(`data`),`update_time`=VALUES(`update_time`)
*/
func TestUpsert1(t *testing.T) {
	report := model.UserMonthlyReport{
		StuId:      2100051684,
		YearMonth:  "202401",
		Version:    "v0.0.1",
		ReadStatus: 1,
		Data:       []byte("{\"key\":\"va1111lue\"}"),
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}
	err := DB.Table("user_monthly_report").Clauses(clause.OnConflict{
		// 这里可以使用数据库唯一键对应的列名或者使用唯一键的名称
		// 数据库 UNIQUE KEY `uni_stu_month` (`stu_id`,`year_month`)
		// 使用这个 Columns:   []clause.Column{{Name: "uni_stu_month"}}, 或者下面这个语句都可以
		Columns:   []clause.Column{{Name: "stu_id"}, {Name: "year_month"}},
		DoUpdates: clause.AssignmentColumns([]string{"data", "update_time"}),
	}).Create(&report).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", report)
}

// 2.通过FirstOrCreate 来实现创建或更新 先查询 然后再根据情况创建或更新，两条sql
// 这个方法实现的upsert不会导致主键不连续，也不依赖唯一索引，更加灵活
/*
	SQL 1: SELECT * FROM `user_monthly_report` WHERE stu_id = 2100051684 and `year_month` = '202410' ORDER BY `user_monthly_report`.`id` LIMIT 1
    查询语句查到数据则进行更新、查不到则进行创建
	SQL 2: UPDATE `user_monthly_report` SET `data`='{"key":"value"}',`update_time`=1724746094 WHERE (stu_id = 2100051684 and `year_month` = '202412') AND `id` = 159
    SQL 3: INSERT INTO `user_monthly_report` (`stu_id`,`year_month`,`version`,`read_status`,`data`,`create_time`,`update_time`) VALUES (2100051684,'202411','v0.0.1',1,'{"key":"va1111lue"}',1724745450,1724745450)
*/
func TestUpsert2(t *testing.T) {
	report := model.UserMonthlyReport{
		StuId:      2100051684,
		YearMonth:  "202412",
		Version:    "v0.0.1",
		ReadStatus: 1,
		Data:       []byte("{\"key\":\"value\"}"),
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}
	// 更新对象需要单独定义，因为FirstOrCreate会将查询到的数据赋值给report这个对象
	update := model.UserMonthlyReport{
		Data:       report.Data,
		UpdateTime: report.UpdateTime,
	}
	err := DB.Table("user_monthly_report").Where("stu_id = ? and `year_month` = ?", report.StuId, report.YearMonth).Assign(&update).FirstOrCreate(&report).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", report)
}
