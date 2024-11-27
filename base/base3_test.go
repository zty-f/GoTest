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

type student struct {
	name string
	age  int
}

func TestStruct(t *testing.T) {
	m := make(map[string]*student)
	stus := []student{
		{name: "pprof.cn", age: 18},
		{name: "测试", age: 23},
		{name: "博客", age: 28},
	}

	for _, stu := range stus { // For range循环的本质是将每个stu对应的内存地址赋值给循环体
		fmt.Printf("%p\n", &stu) // 0x1400001a578 每次都是同一个地址，循环将每个值复制到这个地址
		fmt.Println(&stu)
		m[stu.name] = &stu // 此处是浅拷贝，每次都是将对应的内存地址赋值
	}
	// 循环结束stu地址存储的数据就是最后一次循环的数据
	fmt.Printf("%+v\n", m)
	for k, v := range m {
		fmt.Println(k, "=>", v.name) // 此次打印的都是最后一次循环的数据
	}
	fmt.Println("------------")
	mm := make(map[string]student)
	for _, stu := range stus {
		fmt.Printf("%p\n", &stu)
		mm[stu.name] = stu // 这种方式是深拷贝，每次生成一个新的内存地址赋值
	}
}
