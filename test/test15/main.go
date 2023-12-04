package main

import (
	"fmt"
	"test/test/test11/model"
)

var levelInfo = []model.LevelInfo{
	1: {
		Level:     1,
		LevelName: "初级童生",
		LevelSign: "https://static-inc.xiwang.com/mall/217abc87409474cd022b87cab90a4f94.png",
		LevelValueRange: model.LevelValueRange{
			Begin: 0,
			End:   10000,
		},
	},
	2: {
		Level:     2,
		LevelName: "潜力秀才",
		LevelSign: "https://static-inc.xiwang.com/mall/217abc87409474cd022b87cab90a4f94.png",
		LevelValueRange: model.LevelValueRange{
			Begin: 10000,
			End:   22000,
		},
	},
}

func main() {
	fmt.Println(len(levelInfo))
	fmt.Printf("%+v\n", levelInfo)
	for k, v := range levelInfo {
		fmt.Println(k, v)
	}
	x := map[int]int{}
	delete(x, 3)

}
