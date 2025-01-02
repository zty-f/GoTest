package main

import (
	"fmt"
	"testing"
)

type Test1 struct {
	ID int `gorm:"column:id"`
	A  int `gorm:"column:a"`
	B  int `gorm:"column:b"`
}

func TestFind1(t *testing.T) {
	test := &Test1{}
	db := DB.Table("test1")
	// 不用赋值给db，直接链式调用即可
	db.Where("id = ?", 1)
	db.Where("a = ?", 1)
	err := db.First(test).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", test)
}
