package main

import (
	"fmt"
)

func hello() []string {
	fmt.Println("hello")
	return nil
}

func GetValue() int {
	return 1
}

func test(x interface{}) {
	switch x.(type) {
	case nil:
		fmt.Println("nil")
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("unknown")
	}
}

func main() {
	h := hello
	h()
	if h == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}

	i := GetValue()
	fmt.Println(i)

	test(nil)
	test(1)
	test("1111")
	test(map[string]string{})

	var x = []int{4: 44, 55, 66, 1: 77, 88} //[0 77 88 0 44 55 66]

	fmt.Println(x)

	var s1 []int
	var s2 = []int{}
	if s1 == nil {
		fmt.Println("yes nil")
	} else {
		fmt.Println("no nil")
	}
	if s2 == nil {
		fmt.Println("yes nil")
	} else {
		fmt.Println("no nil")
	}

	A := 65
	a := 97
	fmt.Println(A)
	fmt.Println(string(A))
	fmt.Println(a)
	fmt.Println(string(a))

	m := map[string]string{"a": "b", "b": ""}

	fmt.Println(m["a"])

	v, ok := m["a"]
	fmt.Printf("%v,%v\n", v, ok)
	v, ok = m["b"]
	fmt.Printf("%v,%v\n", v, ok)
	v, ok = m["c"]
	fmt.Printf("%v,%v\n", v, ok)

	arr := []int{1, 2, 3}
	fmt.Println(arr)
	p(arr...) //传入的为原切片，会同步修改，需要注意
	fmt.Println(arr)

	fmt.Println(1 + 333.3)

	//x1 := 1
	//y := 2.3
	//fmt.Println(x1+y)  不能相加
	fmt.Println("------------------")
	tt := [5]int{1, 2, 3, 4, 5}
	t := tt[1:4:4]
	fmt.Println(t)
	fmt.Println("------------------")
	var in interface{}
	fmt.Println(in)
	if in == nil {
		fmt.Println("nil")
	}
	fmt.Println("not nil")
	fmt.Println("------------------")
	s := make(map[string]int)
	/*
		删除 map 不存在的键值对时，不会报错，相当于没有任何作用；获取不存在的减值对时，返回值类型对应的零值，所以返回 0。
	*/
	delete(s, "h")
	fmt.Println(s["h"])
	fmt.Println("------------------")
	i1 := -5
	j1 := +5
	fmt.Printf("%+d %+d\n", i1, j1) //-5 +5

}

func p(num ...int) {
	num[0] = 18
}
