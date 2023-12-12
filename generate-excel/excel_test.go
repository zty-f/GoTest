package main

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/xuri/excelize/v2"
	"log"
	"strconv"
	"testing"
	"time"
)

func TestExcelProcessStream(t *testing.T) {
	var list []OperationLog

	for i := 1; i <= 500000; i++ {
		list = append(list, OperationLog{
			Id:       i,
			UserName: "xxx" + cast.ToString(i),
			IP:       "127.0.0.1" + cast.ToString(i),
			TypeStr:  "xxx" + cast.ToString(i),
			Module:   "xxx" + cast.ToString(i),
			Res:      "xxx" + cast.ToString(i),
			Time:     strconv.FormatInt(time.Now().Unix(), 10),
			Des:      "xxx" + cast.ToString(i),
		})
	}
	begin := time.Now().Unix()
	err := ExcelProcessStream(list).
		Headers("序号", "用户名", "登录IP", "用户类型", "操作模块", "操作结果", "操作时间", "描述").
		Columns("Id", "Username", "IP", "TypeStr", "Module", "Res", "Time", "Des").
		SavePath("demo.xlsx").SetColWidth(2, 7, 25).SetColWidth(8, 8, 80).
		NewStyle(KeyStyle{
			Key:   "style1",
			Style: &excelize.Style{Font: &excelize.Font{Family: "Microsoft YaHei UI", Size: 12}},
		}, KeyStyle{
			Key: "style2",
			Style: &excelize.Style{Font: &excelize.Font{Family: "Microsoft YaHei UI", Size: 12}, Fill: excelize.Fill{
				Type:    "pattern",
				Pattern: 1,
				Color:   []string{"#E8E8E8"},
			}},
		}).Style(func(styleMap map[string]int, x int, y int, val interface{}) (int, error) {
		if y%2 == 0 {
			return styleMap["style2"], nil
		} else {
			return styleMap["style1"], nil
		}
	}).ToExcel().Error
	if err != nil {
		log.Println("err:", err)
	}
	end := time.Now().Unix()
	fmt.Println("表格生成耗费时长：", end-begin, "s")
}
