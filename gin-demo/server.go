package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func Get(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello gin!",
	})
}

type Test struct {
	Name   string `form:"name" json:"name"`
	Brand  string `form:"brand" json:"brand"`
	UserId int    `form:"user_id" json:"user_id"`
}

type Test1 struct {
	Name         string `form:"name" json:"name"`
	Brand        string `form:"brand" json:"brand"`
	UserIdString string `form:"user_id" json:"user_id"`
}

func TestShouldBind(c *gin.Context) {
	test := &Test{}
	err := c.ShouldBindWith(test, binding.Form) //这个不会影响bind次数
	if err != nil {
		fmt.Println(err)
		return
	}
	test2 := &Test{}
	err2 := c.ShouldBind(test2)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	// ShouldBind绑定参数时，如果参数类型是form-data/x-www-form-urlencoded时,可以多次使用ShouldBind
	// 但是如果参数类型是json时，只能使用一次ShouldBind
	//test3 := &Test{}
	//err3 := c.ShouldBindJSON(test3)
	//if err3 != nil {
	//	fmt.Println(err3)
	//	return
	//}
	//test4 := &Test{}
	//err4 := c.ShouldBindJSON(test4)
	//if err4 != nil {
	//	fmt.Println(err4)
	//	return
	//}
	// ShouldBindJSON不能使用多次，尽管是针对不同地址空间的结构体,也不能和shouldBind共同使用多次
	//1.单次解析，追求性能使用 ShouldBindJson，因为多次绑定解析会出现EOF
	//2.多次解析，避免报EOF，使用ShouldBindBodyWith
	c.JSON(200, gin.H{
		"message": "hello TestShouldBind!",
	})
}
