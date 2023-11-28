package main

import (
	"encoding/json"
	"fmt"
	"test/test11/model"
	"time"
)

var user *User

type User struct {
	Name string
}

var levelInfo = map[int]model.LevelInfo{
	1: {
		Level:     1,
		LevelName: "初级童生",
		LevelSign: "等级标志",
		LevelValueRange: model.LevelValueRange{
			Begin: 0,
			End:   10000,
		},
	},
	2: {
		Level:     2,
		LevelName: "潜力秀才",
		LevelSign: "等级标志",
		LevelValueRange: model.LevelValueRange{
			Begin: 10000,
			End:   22000,
		},
	},
	3: {
		Level:     3,
		LevelName: "好学举人",
		LevelSign: "等级标志",
		LevelValueRange: model.LevelValueRange{
			Begin: 22000,
			End:   35000,
		},
	},
	4: {
		Level:     4,
		LevelName: "进阶贡士",
		LevelSign: "等级标志",
		LevelValueRange: model.LevelValueRange{
			Begin: 35000,
			End:   50000,
		},
	},
	5: {
		Level:     5,
		LevelName: "飞跃进士",
		LevelSign: "等级标志",
		LevelValueRange: model.LevelValueRange{
			Begin: 50000,
			End:   80000,
		},
	},
	6: {
		Level:     6,
		LevelName: "探花及第",
		LevelSign: "等级标志",
		LevelValueRange: model.LevelValueRange{
			Begin: 80000,
			End:   120000,
		},
	},
	7: {
		Level:     7,
		LevelName: "超凡榜眼",
		LevelSign: "等级标志",
		LevelValueRange: model.LevelValueRange{
			Begin: 120000,
			End:   200000,
		},
	},
	8: {
		Level:     8,
		LevelName: "天选状元",
		LevelSign: "等级标志",
		LevelValueRange: model.LevelValueRange{
			Begin: 200000,
			End:   300000,
		},
	},
}

func main() {
	fmt.Println(user)
	if user == nil {
		fmt.Println("初始化")
		Init()
	}
	fmt.Println(user)
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().Unix())

	m := make(map[int]string, 0)
	m[7] = "1"
	m[2] = "1"
	m[4] = "1"
	m[5] = "1"

	fmt.Println(len(m))
	fmt.Println(len(levelInfo))

	json.Marshal(m)
}

func Init() {
	user = &User{
		Name: "xwx",
	}
}
