package main

import (
	"encoding/json"
	"fmt"
	"github.com/bytedance/sonic"
	"os"
	"testing"
)

type People struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Tmp  string `json:"tmp"`
}

type MedalDetail struct {
	MedalId    int    `json:"medal_id"`
	MedalType  int    `json:"medal_type"`
	MedalName  string `json:"medal_name"`
	MedalImage string `json:"medal_image"`
	CreateTime int    `json:"create_time"`
}

func TestMarshal(t *testing.T) {
	var tmp People
	// 使用io/ioutil包读取json文件为字符串
	bytes, err := os.ReadFile("/Users/xwx/go/src/test/json/people.json")
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
	fmt.Println(tmp)
	marshal, _ := json.Marshal(tmp)
	fmt.Println(string(marshal))
}

func TestMarshalExchangeDetail(t *testing.T) {
	exchangeDetails := []ExchangeDetail{
		{
			SkuName:    "240601215652100118241212",
			SkuIcon:    "",
			ChangeType: 2,
			PayPoint:   1750,
			CreateTime: 1717250270,
		},
		{
			SkuName:    "240602213402100448601212",
			SkuIcon:    "",
			ChangeType: 2,
			PayPoint:   250,
			CreateTime: 1717335252,
		},
		{
			SkuName:    "启蒙积木军事系列航空母舰拼装玩具",
			SkuIcon:    "https://static0.xesimg.com/udc-o-point/jf/shop/sku/8c1e58c2d8993d68da058a1e3ac872e5.jpg",
			ChangeType: 1,
			PayPoint:   500,
			CreateTime: 1711113251,
		},
		{
			SkuName:    "超精密头盔",
			SkuIcon:    "https://m.xiwang.com/resource/__private__/3650cd180bd6e0042977f35ed53a696d/yasuo/xingji/mao1.png",
			ChangeType: 3,
			PayPoint:   400,
			CreateTime: 1727866297,
		},
		{
			SkuName:    "齐天大圣",
			SkuIcon:    "https://m.xiwang.com/resource/__private__/3650cd180bd6e0042977f35ed53a696d/yasuo/wukong/%E5%85%A81.png",
			ChangeType: 3,
			PayPoint:   3000,
			CreateTime: 1727816297,
		},
	}
	strs := make([]string, 0)
	for _, item := range exchangeDetails {
		marshal, err := sonic.Marshal(item)
		if err != nil {
			fmt.Println(err)
		}
		strs = append(strs, string(marshal))
	}
	bytes, err := sonic.Marshal(strs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bytes))
}

func TestMarshalMedalDetail(t *testing.T) {
	medalDetails := []MedalDetail{
		{
			MedalId:    10070,
			MedalType:  1,
			MedalName:  "毕业成就",
			MedalImage: "https://static-inc.xiwang.com/mall/6cee6f9c5f0ea29fcd86aaf9ae80527e.png",
			CreateTime: 1727866297,
		},
		{
			MedalId:    10062,
			MedalType:  1,
			MedalName:  "全勤战士",
			MedalImage: "https://static-inc.xiwang.com/mall/926524899d7743cf009c8754735ff863.png",
			CreateTime: 1727866397,
		},
		{
			MedalId:    20002,
			MedalType:  2,
			MedalName:  "奋战30天",
			MedalImage: "https://static-inc.xiwang.com/mall/7d8628bd6a54cdab98e824c2fff202cb.png",
			CreateTime: 1727866497,
		},
		{
			MedalId:    40001,
			MedalType:  4,
			MedalName:  "小试牛刀",
			MedalImage: "https://static-inc.xiwang.com/mall/d69d38f7e789abde2a32aa72c645fecb.png",
			CreateTime: 1727866597,
		},
		{
			MedalId:    40007,
			MedalType:  4,
			MedalName:  "连胜14天",
			MedalImage: "https://static-inc.xiwang.com/mall/29e1e9ea7289383861183ef7f2f4ea9a.png",
			CreateTime: 1727866697,
		},
		{
			MedalId:    50002,
			MedalType:  4,
			MedalName:  "钱包守卫者",
			MedalImage: "https://static-inc.xiwang.com/mall/fc69ef81b996dc28f41f8fccb49d9d5d.png",
			CreateTime: 1727866797,
		},
		{
			MedalId:    60001,
			MedalType:  4,
			MedalName:  "坚持打卡7天",
			MedalImage: "https://static-inc.xiwang.com/mall/4b8dddc01e6e96409f4109a5866b30ec.png",
			CreateTime: 1727866897,
		},
		{
			MedalId:    60200,
			MedalType:  4,
			MedalName:  "水星",
			MedalImage: "https://static-inc.xiwang.com/mall/ff3cfdd84fc3a96ce017ec8837b6ed07.png",
			CreateTime: 1727866997,
		},
		{
			MedalId:    70001,
			MedalType:  4,
			MedalName:  "小小书童",
			MedalImage: "https://static-inc.xiwang.com/mall/b79cab2356d4497e607b52b92f8ef82b.png",
			CreateTime: 1727867297,
		},
		{
			MedalId:    10033,
			MedalType:  1,
			MedalName:  "毕业成就",
			MedalImage: "https://static-inc.xiwang.com/mall/3c3636859c95c89ec975cb42ca49f2e2.png",
			CreateTime: 1727868297,
		},
	}
	strs := make([]string, 0)
	for _, item := range medalDetails {
		marshal, err := sonic.Marshal(item)
		if err != nil {
			fmt.Println(err)
		}
		strs = append(strs, string(marshal))
	}
	bytes, err := sonic.Marshal(strs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bytes))
}
