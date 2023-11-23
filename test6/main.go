package main

import "fmt"

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")

}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func main() {
	t := Teacher{}
	t.ShowA()

	list := new([]int) //返回的是一个指针变量
	*list = append(*list, 1, 2, 3, 4)
	fmt.Println(*list)

	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	fmt.Println(sn1 == sn2)

	sn3 := struct {
		name string
		age  int
	}{name: "qq", age: 11}

	fmt.Println(sn3)
	// 结构体的比较需要两个结构体的字段类型和顺序都相同，并且每个字段都是可比较的类型
	//fmt.Println(sn1 == sn3)  // 编译报错

	p := &sn3
	fmt.Printf("%p\n", p)
	fmt.Printf("%T\n", p)

	fmt.Println(p.name)
	fmt.Println((*p).name)

	str1 := 'a' + '1'
	fmt.Println(str1) //146 字符都转为ASCLL码  97 + 49
	str2 := "abc" + "123"
	fmt.Println(str2)
	//str3 := '1' + "333"  // ''单引号里面只有一个字符可以相加，两个字符相加需要用双引号
	//str4 := '12' + "333"
	println(fmt.Sprintf("abc%d", 123))

	//xx := 'qqq' // 编译报错，因为'qqq'是字符串需要使用双引号
}
