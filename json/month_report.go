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
	SkuName    string `json:"sku_name"`    // 商品名称（购课：订单号）
	SkuIcon    string `json:"sku_icon"`    // 商品图链接
	ChangeType int    `json:"change_type"` // 1:商品；2:金币抵现 3：装扮兑换
	PayPoint   int    `json:"pay_point"`   // 兑换金币数量
	CreateTime int    `json:"create_time"` // 这笔兑换产生的时间（用于排序）
}

type Test struct {
	A string `json:"希望学-PK宝箱"`
}

func main() {
	var tmp MonthlyReportMessage
	// 使用io/ioutil包读取json文件为字符串
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
