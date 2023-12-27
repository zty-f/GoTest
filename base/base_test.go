package base

import (
	"fmt"
	"testing"
)

func Test_m1(t *testing.T) {
	m1()
}

type A interface {
	Show()
}

type B struct{}

func (stu *B) Show() {
	fmt.Println("show")
}

func Test1(t *testing.T) {
	var s *B
	if s == nil {
		fmt.Println("s is nil")
	} else {
		fmt.Println("s is not nil")
	}
	var p A = s
	if p == nil {
		fmt.Println("p is nil")
	} else {
		fmt.Println("p is not nil")
		fmt.Println(p)
		fmt.Printf("%T\n", p)
		p.Show()
	}
}

type User struct {
	Name string
	Age  int
}

func Test2(t *testing.T) {
	var arr []*User
	// 以上创建方式不能直接赋值  arr[0] = ..  可以使用append 会自动扩容
	arr = append(arr, &User{
		Name: "zhangsan",
		Age:  18,
	})
	fmt.Printf("%+v\n", arr[0])

	var arr2 = make([]User, 2, 2)
	// 以上创建方式同样也不能直接赋值  arr[0] = ..  make创建的第一个参数为长度  第二个参数为容量 指定长度内可以直接赋值  容量可以使用append的时候不会频繁扩容
	arr2[1] = User{
		Name: "lisi",
		Age:  18,
	}
	fmt.Printf("%+v\n", arr2)
}
