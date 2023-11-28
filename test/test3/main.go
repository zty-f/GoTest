package main

import (
	"fmt"
	"reflect"
)

type A struct {
	X int
	Y string
	Z int
}

func main() {
	a := A{
		X: 1,
		Y: "222",
		Z: 3,
	}
	b := A{
		X: 3,
	}
	va := reflect.ValueOf(a)
	vb := reflect.ValueOf(b)
	num := va.NumField()
	fmt.Println(num)

	for i := 0; i < num; i++ {
		fmt.Println(reflect.TypeOf(a).Field(i).Name)
		fmt.Println(va.Field(i).String())
		fmt.Println(reflect.TypeOf(b).Field(i).Name)
		fmt.Println(vb.Field(i).String())
	}
}
