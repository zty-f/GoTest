package base

import (
	"fmt"
	"testing"
)

func TestPtr(t *testing.T) {
	a := 1
	b := &a
	fmt.Println(b)
	fmt.Println(*b)
	fmt.Println("--------------")
	x := new(int)
	fmt.Println(x)
	fmt.Println(*x)
	fmt.Println("--------------")
	y := new(map[int]int)
	fmt.Println(y)
	fmt.Println(*y)
	// (*y)[0] = 1 // 会报错panic
	fmt.Println((*y)[0])
	fmt.Println("--------------")
	z := make(map[int]int)
	fmt.Println(z)
	z[1] = 4
	fmt.Println(z[1])
	fmt.Println("--------------")
	// make不能用于类型的初始化，但是可以用于slice、map以及channel的初始化
	// c := make(string,10)
	// c := make(int,10)
	// c := make(bool,10)
	/*
	   1.二者都是用来做内存分配的。
	   2.make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
	   3.而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。

	   new适用于普通类型的初始化以及内存分配
	   make只适用于slice、map以及channel的初始化
	*/
}
