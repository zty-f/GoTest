package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func test() (a, b int) {
	return 1, 2
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Like Like   `json:"like"`
}

type Like struct {
	A string `json:"a"`
}

type Data struct {
	User []User `json:"userList"`
}

func main() {
	var a int
	//a, _ := test() 已经声明了a，不能使用:=
	a, b := test()
	fmt.Println(a, b) //当多值赋值时，:= 左边的变量无论声明与否都可以
	jsonFile, err := os.Open("test9/userList.json")
	if err != nil {
		fmt.Println("error opening json file", err.Error())
		return
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("error reading json file")
		return
	}
	var user Data
	json.Unmarshal(jsonData, &user)
	fmt.Printf("%+v\n", user)
	x, _ := json.Marshal(user)
	fmt.Println(string(x))
}
