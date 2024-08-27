package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"test/gorm/model"
	"testing"
	"time"
)

var source = `{
    "stuId": 2100051684,
    "growthValue": 84934,
    "level": 3,
    "levelName": "好学举人",
    "star": 3,
    
    "year": "2024",
    "month": "2",
    "beginDate": "02.01",
    "endDate": "02.29",
    "keyword": "完课标兵",

    "powerValue":800,
    "taskFinishList":[
      {
        "taskId":2001200,
        "taskName":"课前预习",
        "taskFinishNum":3
      }
    ],
    
    "monthGrowthValue": 2108,
    "growthValueOverRate": 20,
    "levelDesc": "升级到",
    "nextLevel":4,
    "toNextLevelGrowthValue":234,
    "levelPercent": 60,
    "valueChangeList":[
      {
        "dateRange": "6.1-6.7",
        "addValue": 234
      },
      {
        "dateRange": "6.8-6.15",
        "addValue": 233
      },
      {
        "dateRange": "6.16-6.24",
        "addValue": 23
      },
      {
        "dateRange": "6.25-6.30",
        "addValue": 124
      }
    ],
    
    "goldIncomeNum": 4564,
    "classGoldIncomeNum": 5808,
    "classGoldIncomeRate": 45,
    "levelGoldIncomeNum": 1259,
    "levelGoldIncomeRate": 25,
    "otherGoldIncomeNum": 345,
    "otherGoldIncomeRate": 30,
    "goldExpendNum": 39000,
    "exchangeSkuNum": 5,
    "exchangeSkuList": [
      {
        "skuId": "",
        "skuName": "主讲课堂互动灯牌-星耀粉",
        "skuIcon": "https://static0.xesimg.com/udc-o-point/jf/shop/sku/e63ab729eb3a83b3f03fc9a8f6863441.png",
        "payPoint": 1299,
        "createTime": 1711260203
      },
      { 
        "skuId": "26363773327637",
        "skuName": "金币抵现",
        "skuIcon": "https://static0.xesimg.com/udc-o-point/jf/shop/sku/154f64a9275fd2b8190237ba9dcdf4df.jpg",
        "payPoint": 7299,
        "createTime": 1710567469
      }
    ]
  }`

func TestMonthReportMsgUnmarshal(t *testing.T) {
	var data model.MonthlyReportResponse

	err := json.Unmarshal([]byte(source), &data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", data)
}

func AddPrefix0(str string) string {
	if cast.ToInt(str) < 10 {
		return "0" + str
	}
	return str
}

func TestAdd(t *testing.T) {
	var data model.MonthlyReportResponse

	err := json.Unmarshal([]byte(source), &data)
	if err != nil {
		fmt.Println(err)
		return
	}
	bytes, _ := json.Marshal(data)
	report := model.UserMonthlyReport{
		StuId:      data.StuId,
		YearMonth:  data.Year + AddPrefix0(data.Month),
		Version:    "v0.0.1",
		Data:       []byte(string(bytes)),
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}
	err = DB.Table("user_month_report").Create(&report).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", report)
}

func TestGet(t *testing.T) {
	var report model.UserMonthlyReport
	err := DB.Table("user_month_report").Where("stu_id = ? and `year_month` = ?", 2100051684, "202402").First(&report).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", report)

	var data model.MonthlyReportResponse
	err = json.Unmarshal([]byte(report.Data), &data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", data)
}
