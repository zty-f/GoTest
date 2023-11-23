package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"test/gorm/model"
)

var DB *gorm.DB

func init() {
	// 初始化
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "xw_mall_dev_rw:8Ufp%mUzg*bnkHMZVt8MwVdCdV&PzG@tcp(rm-2ze58oz13gl4089i37o.mysql.rds.aliyuncs.com:3306)/growth_center?charset=utf8&loc=Local&parseTime=True", // DSN data source name
		DefaultStringSize:         256,                                                                                                                                                         // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                                                                                                                                                        // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                                                                                                                                                        // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                                                                                                                                        // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                                                                                                                                                       // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	DB = db
}

func main() {
	testUpdate()
}

func testUpdate() {
	var usertask = model.UserTask{
		CreateTime: 11111122,
		UpdateTime: 0,
		IsDeleted:  0,
		Extra:      "",
	}

	t := DB.Debug().Model(&usertask).Where("unique_id = ?", "11111").Updates(&usertask)

	fmt.Printf("%+v\n", t)
	fmt.Printf("%+v\n", usertask)
}
