package base

import (
	"encoding/json"
	"fmt"
	"log"
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

func TestSliceAndMapInit(t *testing.T) {
	// 切片为初始化可以进行操作，加值，但是不能取值，会报错
	var s []int
	// fmt.Println(s[0])  // panic: runtime error: index out of range
	s = append(s, 1)
	fmt.Println(s)
	fmt.Println("------------")
	// make 切片会给每个元素设置0值，append是往后面加值
	s1 := make([]int, 5)
	s1 = append(s1, 99)
	fmt.Println(s1) // [0 0 0 0 0 99]
	fmt.Println("------------")

	var m map[int]int
	// m[1] = 2 // panic: assignment to entry in nil map  map为nil时不能进行操作 需要初始化分配空间
	fmt.Println(m)
	// 未初始化的数组可以进行取值操作，默认为字段零值
	fmt.Println(m[1]) // 0

	m = nil

	fmt.Println(m)
	fmt.Println(m[1])
}

func TestStringAndRune(t *testing.T) {
	s := "abcdefg"
	fmt.Println(string(s[0]))
	fmt.Println(s)
	sb := []byte(s)
	fmt.Println(string(sb[0]))
	fmt.Println(sb)

	fmt.Println("------------")

	// 中文字符的时候转换成byte数组的时候字符不能一一对应，string的底层默认是不可修改的[]byte
	s = "你好呀我是谁"
	fmt.Println(string(s[8]))
	fmt.Println(s)
	sb = []byte(s)
	fmt.Println(string(sb[8]))
	fmt.Println(sb)

	fmt.Println("------------")
	s = "你好呀我是谁"
	fmt.Println(string(s[8]))
	fmt.Println(s[0:2])
	// rune是int32的别名,使用这个数组转换能够是中文字符一一对应
	sru := []rune(s)
	fmt.Println(string(sru[5]))
	fmt.Println(sru[0:2])
}

func isChar(x byte) bool {
	// switch 语句中的 case 代码块会默认带上 break，但可以使用 fallthrough 来强制执行下一个 case 代码块,只会穿透到下一个 case
	switch x {
	case ' ':
		fmt.Println("空格", x)
		fallthrough // 返回 true
	case '\t':
		fmt.Println("制表符", x)
		return true
	case '\n':
		fmt.Println("换行", x)
		return true
	}
	fmt.Println("switch end")
	return false
}

func TestSwitch(t *testing.T) {
	fmt.Println(isChar(' '))
	fmt.Println("------------")
	fmt.Println(isChar('\t'))
}

func TestMarshalInt(t *testing.T) {
	var data = []byte(`{"status": 200}`)
	var result map[string]interface{}

	if err := json.Unmarshal(data, &result); err != nil {
		log.Fatalln(err)
	}
	// 在 encode/decode JSON 数据时，Go 默认会将数值当做 float64 处理。
	fmt.Printf("%T\n", result["status"])
	fmt.Println("------------")
	var res1 map[string]int
	if err := json.Unmarshal(data, &res1); err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%T\n", res1["status"])
}

func TestPanicRecover1(t *testing.T) {
	// 错误的 recover 调用示例
	recover()         // 什么都不会捕捉
	panic("not good") // 发生 panic，主程序退出
	recover()         // 不会被执行
	println("ok")
}

func TestPanicRecover2(t *testing.T) {
	// 错误的 recover 调用示例
	defer recover()
	panic("not good")
	fmt.Println("ok")
}

func TestPanicRecover3(t *testing.T) {
	// defer 调用时多层嵌套依然无效。
	defer func() {
		func() {
			recover()
		}()
	}()
	panic(1)
}

func TestPanicRecover4(t *testing.T) {
	// 正确的 recover 调用示例 必须在 defer 函数中直接调用才有效。
	defer func() {
		fmt.Println("recovered: ", recover())
	}()
	panic("not good")
	fmt.Println("ok")
}
