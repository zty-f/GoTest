package main

import "fmt"

var m = map[int]string{
	1:  "1",
	2:  "2",
	3:  "3",
	4:  "4",
	5:  "5",
	6:  "6",
	7:  "7",
	8:  "8",
	9:  "9",
	10: "10",
}
var arr = []string{
	1: "1",
	2: "2",
	3: "3",
	4: "4",
	5: "5",
	6: "6",
	7: "7",
	8: "8",
	9: "9",
}

func TraverseMap() {
	for k, v := range m {
		_ = k
		_ = v
	}
}

func TraverseArray() {
	for k, v := range arr {
		_ = k
		_ = v
	}
}

func main() {

	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, val := range slice {
		fmt.Printf("%p\n", &val)
		m[key] = &val
	}

	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
}
