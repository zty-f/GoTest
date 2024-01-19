package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"testing"
)

func TestCreate(t *testing.T) {
	// 连接到 MySQL 数据库
	dsn := "root:123456@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 自动迁移表结构
	tablePrefix := "user"
	numTables := 5
	sql := "CREATE TABLE %s (id INT PRIMARY KEY, name VARCHAR(255))"

	for i := 1; i <= numTables; i++ {
		tableName := fmt.Sprintf("%s_%d", tablePrefix, i)
		err := db.Exec(fmt.Sprintf(sql, tableName)).Error
		if err != nil {
			panic(err)
		}
		fmt.Printf("Created table: %s\n", tableName)
	}

	// 关闭数据库连接
	dbSQL, err := db.DB()
	if err != nil {
		panic(err)
	}
	dbSQL.Close()
}

func TestGenerateSql(t *testing.T) {

	// 自动迁移表结构
	tablePrefix := "user"
	numTables := 5
	sql := "CREATE TABLE `%s` (\n`id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',\n`stu_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',\n`value` int(11) NOT NULL DEFAULT '0' COMMENT '成长值，可正可负',\n`source` int(11) NOT NULL DEFAULT '-1' COMMENT '成长值更新来源， 10:初始化, 20: 任务 30:激活，',\n`task_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '该成长值记录对应的任务id',\n`user_task_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '关联user_task表，对应用户任务记录id\\n',\n`value_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '获得成长值的时间，要和create_time区分',\n`create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',\n`ext_data` text NOT NULL COMMENT '扩展字段，json格式',\nPRIMARY KEY (`id`) USING BTREE,\nKEY `idx_uid_src` (`stu_id`,`source`),\nKEY `idx_uid_vtime` (`stu_id`,`value_time`),\nKEY `idx_utaskid` (`user_task_id`)\n) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='成长值记录表';"
	result := ""
	for i := 1; i <= numTables; i++ {
		tableName := fmt.Sprintf("%s_%d", tablePrefix, i)
		x := fmt.Sprintf(sql, tableName)
		result += x + "\n"
	}

	// 将result字符串写入文件
	err := os.WriteFile("test.sql", []byte(result), 0644)
	if err != nil {
		panic(err)
	}
}
