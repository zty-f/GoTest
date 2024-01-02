package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name int  `json:"name,int:string"`
	Age  bool `json:"age,int:string"`
}

func main() {
	//t := util.Now().Unix() * 100000000000
	//fmt.Println(util.Now().Unix() * 1000)
	//user := User{
	//	Name: int(t),
	//	Age:  true,
	//}
	//
	//jsonStr, _ := json.Marshal(user)
	jsonS := "{\"name\":\"4119396136614035456\",\"age\":\"true\"}" //"{"name":"4119396136614035456","age":"true"}

	fmt.Println(string(jsonS))

	var user2 User
	json.Unmarshal([]byte(jsonS), &user2)

	fmt.Println(user2)

	fmt.Printf("%T\n", user2.Name)
	fmt.Printf("%T\n", user2.Age)

	//var x int = -1
	//
	//fmt.Println(-1)
	//fmt.Println(x)
}
