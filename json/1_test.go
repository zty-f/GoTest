package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

type People struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Tmp  string `json:"tmp"`
}

func TestMarshal(t *testing.T) {
	var tmp People
	//使用io/ioutil包读取json文件为字符串
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
