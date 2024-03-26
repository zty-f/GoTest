package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type MonthlyReportMessage struct {
	SceneGoldRate string `json:"scene_gold_rate"`
}

type ExchangeDetail struct {
	SkuName    string `json:"sku_name"`
	SkuIcon    string `json:"sku_icon"`
	PayPoint   int    `json:"pay_point"`
	CreateTime int    `json:"create_time"`
}

type Test struct {
	A string `json:"希望学-PK宝箱"`
}

func main() {
	var tmp MonthlyReportMessage
	//使用io/ioutil包读取json文件为字符串
	bytes, err := os.ReadFile("json/month_report.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(string(bytes))
	err = json.Unmarshal(bytes, &tmp)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var t Test
	err = json.Unmarshal([]byte(tmp.SceneGoldRate), &t)
	fmt.Println(tmp)
	fmt.Println(t)
}
