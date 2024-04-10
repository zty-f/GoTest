package base

import (
	"codeup.aliyun.com/61e54b0e0bb300d827e1ae27/backend/golib/logger"
	"context"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"sort"
	"strings"
	"testing"
	"time"
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

func TestBinary(t *testing.T) {
	fmt.Printf("%b\n", 1)
	fmt.Printf("%b\n", int64(-5))
	fmt.Printf("%b\n", 1<<4)
	fmt.Printf("%b\n", 1<<4-1)
}

type R struct {
	Include bool
	Value   int
}

// 数值排序测试
func TestRank(t *testing.T) {
	arr := []R{
		{
			Include: true,
			Value:   1,
		},
		{
			Include: false,
			Value:   2,
		},
		{
			Include: true,
			Value:   3,
		},
		{
			Include: false,
			Value:   4,
		},
		{
			Include: true,
			Value:   5,
		},
	}
	// 排序 现根据Include进行排序，Include为true的在前面，false的在后面，Include值相同的按照Value大小进行排序
	sort.Slice(arr, func(i, j int) bool {
		if arr[i].Include == arr[j].Include {
			return arr[i].Value > arr[j].Value // 大于表示从大到小排序
		}
		return arr[i].Include
	})
	fmt.Println(arr)
}

func TestFormat(t *testing.T) {
	jumpInfo := "1、购买系统课/续报时，金币可以作为现金使用\n2、250金币抵现1元\n3、每个订单最多能抵扣%d元\n4、抵扣现时，可以使用的金币数量是250的整数倍。"
	s := fmt.Sprintf(jumpInfo, 200)
	fmt.Println(s)

	b := false
	if !b {
		fmt.Println("true")
	}
	fmt.Println(b)
}

func TestByte2String(t *testing.T) {
	x := []byte{0x7b, 0x22, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x35, 0x30, 0x30, 0x31, 0x31, 0x30, 0x30, 0x2c, 0x22, 0x73, 0x74, 0x75, 0x5f, 0x69, 0x64, 0x22, 0x3a, 0x20, 0x32, 0x34, 0x34, 0x32, 0x36, 0x36, 0x39, 0x2c, 0x22, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x3a, 0x22, 0x61, 0x62, 0x63, 0x22, 0x2c, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x22, 0x2c, 0x22, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x3a, 0x31, 0x37, 0x30, 0x34, 0x37, 0x30, 0x31, 0x36, 0x35, 0x39, 0x7d}
	fmt.Println(string(x))
}

func f(n int) (r int) {
	defer func() {
		r += n
		err := recover()
		fmt.Println(err)
	}()
	var f func()
	defer f() // 此时f为空，会直接panic
	f = func() {
		r += 2
	}
	return n + 1
}
func TestDefer(t *testing.T) {
	//fmt.Println(f(3))
	//var m map[int]int
	//fmt.Println(m[1])
	//fmt.Println(m[2])
	s1 := []int{1, 2, 3}
	s2 := s1[1:]
	s2[1] = 4
	fmt.Println(s1)
	s2 = append(s2, 5, 6, 7)
	s1 = append(s1, 8, 9, 10)
	fmt.Println(s1)
	fmt.Println(s2)
	if a := 1; false {
	} else if b := 2; false {
	} else {
		println(a, b)
	}
}

func TestA(t *testing.T) {
	var peo People = &Student{}
	think := "speak"
	fmt.Println(peo.Speak(think))
}

func TestExtendInterface(t *testing.T) {
	var peo People = &Student{}
	think := "speak"
	fmt.Println(peo.Speak(think))
	fmt.Println(math.MaxInt64)
}

func TestEqual(t *testing.T) {
	a := [2]int{5, 6}
	b := [3]int{5, 6}
	fmt.Println(a) // [5 6]
	fmt.Println(b) // [5 6 0]
}

func TestContext(t *testing.T) {
	var ctx = context.Background()
	ctx = context.WithValue(ctx, "key", "value")
	key := logger.ContextKey("x_trace_id")
	i := 0
	for {
		// 当ctx这样嵌套赋值的时候，假如要获取最开始的赋值的那个key，耗时就会越来越高，解决办法就是每次新建一个变量：ctx := context.WithValue(ctx, key, i)
		ctx = context.WithValue(ctx, key, i)
		nano := time.Now().UnixMicro()
		fmt.Println(ctx.Value("key"), time.Now().UnixMicro()-nano)
		i++
	}
}

func TestDb(t *testing.T) {
	x := 2100060321
	fmt.Println(x % 100)
}

func TestSlice(t *testing.T) {
	s := make([]int, 0)
	for i := 0; i < 25; i++ {
		s = append(s, i)
	}
	length := len(s)
	for i := 0; i < length; i += 20 {
		end := i + 20
		if end > length {
			end = length
		}
		fmt.Println(s[i:end])
	}
}

func TestSplit(t *testing.T) {
	stuCourseId := strings.Split("2100051764-2001400-15801423-1419403-100", "-")[2]
	fmt.Println(stuCourseId)
}

var DataExistErr = errors.New("data exists")

func TestErrEqual(t *testing.T) {
	if errors.Is(getErr(), DataExistErr) {
		fmt.Println("data exists")
	} else {
		fmt.Println("data not exists")
	}
}

func getErr() error {
	return DataExistErr
}

func TestHappyNewYear(t *testing.T) {
	fmt.Println("工作顺利~")
	fmt.Println("happy new year")
}

func app() func(string) string {
	t := "Hi"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}

func TestClosure(t *testing.T) {
	a := app()
	b := app()
	a("go")
	fmt.Println(b("All"))
	fmt.Println(a("All"))
}

type S struct {
	A int
	B int
}

func TestStableSort(t *testing.T) {
	arr := []S{
		{A: 1, B: 1},
		{A: 2, B: 2},
		{A: 3, B: 1},
		{A: 4, B: 3},
		{A: 5, B: 1},
		{A: 6, B: 2},
		{A: 7, B: 3},
		{A: 8, B: 1},
		{A: 9, B: 3},
	}
	// 稳定排序，假如数组原本就根据A有序，后续根据B字段进行排序不会打乱原有的排序
	fmt.Println(arr)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].B < arr[j].B
	})
	fmt.Println(arr)
}

func TestFmt(t *testing.T) {
	n, err := fmt.Printf("a %d", 1)
	fmt.Println(n, err)
}

func TestSliceSite(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := s1[1:]
	s2[1] = 4
	fmt.Println(s1)
	fmt.Printf("%p\n", &s1[2]) // 0x1400001e340
	fmt.Printf("%p\n", &s2[1]) // 0x1400001e340
	s2 = append(s2, 5, 6, 7)   // 切片扩容底层会生成新的数组，不会影响原数组的切片
	fmt.Printf("%p\n", &s1[2]) // 0x1400001e340
	fmt.Printf("%p\n", &s2[1]) // 0x1400001a578
	fmt.Println(s1)
	fmt.Println(s2)
}

func TestForRange(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	for i, v := range arr {
		fmt.Println(i, v)
	}
	for i := range arr {
		fmt.Println(i)
	}
}

func ForRange() {
	arr := []int{1, 2, 3, 4, 5}
	for i, _ := range arr {
		_ = i
	}
}

func ForI() {
	arr := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		_ = i
	}
}

func BenchmarkForRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ForRange()
	}
}

func BenchmarkForI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ForI()
	}
}

func TestCalc(t *testing.T) {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
	/*
		10 1 2 3
		20 0 2 2
		2 0 2 2
		1 1 3 4
	*/
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

// 生成一个可以指定长度的随机字符串包含大小写字母以及阿拉伯数字
func GetRandomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(65 + rand.Intn(25)) //A=65
	}
	return string(bytes)
}

func TestGetRandomString(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(GetRandomString(10))
	}
}

func TestMap(t *testing.T) {
	m := map[int]int{1: 2, 2: 3}
	fmt.Println(m[1])
	fmt.Println(m[2])
	fmt.Println(m[3])
	fmt.Println(m[0]) // 不存在的key默认值为0
}

func Test取余(t *testing.T) {
	fmt.Println(5 % 2)
	fmt.Println(6 % 2)
	fmt.Println(7 % 2)
	fmt.Println(8 % 2)
	fmt.Println(9 % 2)
}

func TestTime(t *testing.T) {
	// 使用 AddDate 的时候需要格外注意  只要把进行增减的这个时间设置为一号就可以避免如下的问题
	date := time.Date(2024, 1, 31, 0, 0, 0, 0, time.Local)
	format := date.Format("2006-01-02-15-04-05")
	fmt.Println(format)
	date = date.AddDate(0, 1, 0)
	format = date.Format("2006-01-02-15-04-05")
	fmt.Println(date)
	date = time.Date(2024, 3, 31, 0, 0, 0, 0, time.Local)
	format = date.Format("2006-01-02-15-04-05")
	fmt.Println(format)
	//date = date.AddDate(0, -1, 0)
	//format = date.Format("2006-01-02-15-04-05")
	//fmt.Println(date)
	date = time.Now().Add(-time.Hour * 24 * 30)
	format = date.Format("2006-01-02-15-04-05")
	fmt.Println(date)
}

func TestMap2(t *testing.T) {
	m := map[int]int{1: 2, 2: 3}
	fmt.Println(m[1])
	fmt.Println(m[2])
	fmt.Println(m[3])
	fmt.Println(m[0]) // 不存在的key默认值为0
	if v, ok := m[1]; ok {
		fmt.Println(v, ok)
	}
}

func TestSlice11(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(s[0:4])
}
