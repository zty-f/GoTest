package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type User struct {
	Name int  `json:"name,string"`
	Age  bool `json:"age,string"`
}

func main() {
	t := time.Now().Unix() * 100000000000
	fmt.Println(time.Now().Unix() * 1000)
	user := User{
		Name: int(t),
		Age:  true,
	}

	jsonStr, _ := json.Marshal(user)

	fmt.Println(string(jsonStr))

	var user2 User
	json.Unmarshal(jsonStr, &user2)

	fmt.Println(user2)

	fmt.Printf("%T\n", user2.Name)
	fmt.Printf("%T\n", user2.Age)

	var x int = -1

	fmt.Println(-1)
	fmt.Println(x)
}
