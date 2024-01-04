package base

import (
	"fmt"
	"math"
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

func Test3(t *testing.T) {
	var user = &User{
		Name: "zhangsan",
		Age:  18,
	}
	fmt.Println(user.Age)
	fmt.Println((*user).Age)
	fmt.Printf("%T\n", user)
	fmt.Printf("%T\n", *user)
	fmt.Printf("%+v\n", user)
	fmt.Printf("%+v\n", *user)
}

func (u *User) A() {
	fmt.Println("A")
}

func (u User) B() {
	fmt.Println("B")
}

func Test4(t *testing.T) {
	var user1 = &User{}
	user1.A()
	user1.B()
	var user2 = User{}
	user2.A()
	user2.B()
}

type Person struct {
	Name string
}

var list map[string]Person

func Test5(t *testing.T) {

	list = make(map[string]Person) //不需要指定大小和容量，会自动扩容

	student := Person{"Aceld"}

	list["student"] = student
	// 下列代码不能直接进行赋值操作，是值引用，只读
	//list["student"].Name = "Aceld2"

	fmt.Println(list["student"])
	fmt.Println(len(list))
}

func Test6(t *testing.T) {
	fmt.Println(math.MaxInt64)
	fmt.Println(math.MaxInt32)
}

func Test7(t *testing.T) {
	var m = make(map[int]interface{})
	m[1] = "1"
	fmt.Println(m[1])
	//获取不存在的map键值不会报错,会返回零值
	fmt.Println(m[2])
	fmt.Println(m[3])
	m[3] = "3"
	fmt.Println(m[3])

	fmt.Println(&m)
	fmt.Printf("%p\n", m)
	fmt.Printf("%p\n", &m)
	testMap(m)
}

// 引用地址的传递只会生成一个引用副本，指向的数据区域的内存地址不会发生变化
func testMap(m map[int]interface{}) {
	fmt.Println(&m)
	fmt.Printf("%p\n", m)
	fmt.Printf("%p\n", &m)
}
