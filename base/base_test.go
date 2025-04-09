package base

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/cast"
	"math"
	"math/rand"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"sync"
	"test/util"
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

	list = make(map[string]Person) // 不需要指定大小和容量，会自动扩容

	student := Person{"Aceld"}

	list["student"] = student
	// 下列代码不能直接进行赋值操作，是值引用，只读
	// list["student"].Name = "Aceld2"

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
	// 获取不存在的map键值不会报错,会返回零值
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
func TestDefer1(t *testing.T) {
	// fmt.Println(f(3))
	// var m map[int]int
	// fmt.Println(m[1])
	// fmt.Println(m[2])
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
	key := ctx.Value("x_trace_id")
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
		bytes[i] = byte(65 + rand.Intn(25)) // A=65
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
	// date = date.AddDate(0, -1, 0)
	// format = date.Format("2006-01-02-15-04-05")
	// fmt.Println(date)
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

func TestIndex(t *testing.T) {
	sourceData := "{112233gold_price:1}"
	if strings.Index(sourceData, "gold_price") == -1 {
		fmt.Println("不存在")
		return
	}
	fmt.Println("存在")
	fmt.Println(sourceData)
}

func TestStringSlice(t *testing.T) {
	var res []string
	c := []string{"3vjzP3", "CqSr40", "VfI63P"}
	b, _ := json.Marshal(c)
	str := string(b)
	err := json.Unmarshal([]byte(str), &res)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

// 数组转换为字符串
func TestStrings2String(t *testing.T) {
	c := []string{"3vjzP3", "CqSr40", "VfI63P"}
	fmt.Println(strings.Join(c, ","))
}

func TestSplit11(t *testing.T) {
	res := make([]string, 0)
	x := "1,2,3,4,5"
	res = append(res, strings.Split(x, ",")...)
	y := ""
	res = append(res, strings.Split(y, ",")...) // 为空的也会生成一个append，需要特殊处理才行
	fmt.Println(res)                            // [1 2 3 4 5 ]
	fmt.Println(len(res))                       // 6
}

func TestTimeRange(t *testing.T) {
	planStartTime, _ := util.StrToTime("2024-05-31 19:00:00")
	dayTime := util.GetDayStartTime(planStartTime)
	fmt.Println(dayTime)
	before7DayTime := dayTime.Add(-1 * 7 * 24 * time.Hour)
	fmt.Println(before7DayTime)
}

func TestMd5(t *testing.T) {
	str := cast.ToString(76232298) + "edc_exam_token_string"
	token := md5.Sum([]byte(str))
	fmt.Printf("%x\n", token)
	fmt.Println(hex.EncodeToString(token[:]))
}

func Test修改用户类型(t *testing.T) {
	stuId := 2100051684
	fmt.Println((stuId / 8) % 3)
	fmt.Println(stuId % 8)
}

func NumberIntegralMultipleCountInRange(number int, start int, end int) int {
	if number <= 0 {
		return 0
	}

	if end <= start {
		return 0
	}

	if end < number {
		return 0
	}

	counts := 0
	cursor := start
	for {
		if cursor > start && cursor <= end {
			counts = counts + 1
		}
		if cursor >= end {
			break
		}

		cursor = cursor + number
	}

	return counts
}

func TestNumberIntegralMultipleCountInRange(t *testing.T) {
	fmt.Println(NumberIntegralMultipleCountInRange(10, 40, 60))
}

func TestRoundDown(t *testing.T) {
	fmt.Println(RoundDown(11))
	fmt.Println(RoundDown(20))
	fmt.Println(RoundDown(34))
	fmt.Println(RoundDown(45))
	fmt.Println(RoundDown(56))
	fmt.Println(RoundDown(101))
	fmt.Println(RoundDown(100))
	fmt.Println(RoundDown(203))
}

func RoundDown(number int) int {
	return number - number%10
}

// 错题代码实例，极其不推荐
func WrongExample() {
	result := make([]string, 0)
	for i := 0; i < 10; i++ {
		res := "通过rpc或者http的方式请求到数据" // 平均一次请求100ms，10次请求就是1s，假如需要循环100次，那就是10s，这样就会导致接口严重超时
		result = append(result, res)
	}
	fmt.Println(result)
	/*
		解决办法：
		 1. 通过批量请求的方式来减低每次网络请求造成的耗时问题，一次性把数据全部请求出来进行处理
		 2. 通过并发的方式来减少每次网络请求造成的耗时问题，一次性把数据全部请求出来进行处理
		 3. 如果对于列表信息请求接口返回中不是用户这次请求需要展示的数据，而是类似于模块的跳转链接的这类情况，可以不用在列表接口里
			提前把数据构造好，除非数据都是很好构造的情况，不需要额外请求数据，其余情况都可以在用户真实点击的时候再去请求接口获取跳转链接，把过程拆分成多个阶段。
	*/
}

func TestRandom(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(10))
	fmt.Println(rand.Intn(10))
	fmt.Println(rand.Intn(10))
	fmt.Println(rand.Intn(10))
	fmt.Println(rand.Intn(10))
	fmt.Println(rand.Intn(10))
	fmt.Println(rand.Intn(10))
	// 生成一个随机数 范围在1 - 100
	fmt.Println(rand.Intn(100) + 1)
}

func trimPrefix0(str string) string {
	return strings.TrimPrefix(str, "0")
}

func TestString1(t *testing.T) {
	fmt.Println(trimPrefix0("01"))
	fmt.Println(trimPrefix0("08"))
	fmt.Println(trimPrefix0("09"))
	fmt.Println(trimPrefix0("10"))
	fmt.Println(trimPrefix0("11"))
	fmt.Println(trimPrefix0("12"))
}

func AddPrefix0(str string) string {
	if strings.HasPrefix(str, "0") {
		return str
	}
	if cast.ToInt(str) < 10 {
		return "0" + str
	}
	return str
}

func TestString2(t *testing.T) {
	fmt.Println(AddPrefix0("1"))
	fmt.Println(AddPrefix0("2"))
	fmt.Println(AddPrefix0("3"))
	fmt.Println(AddPrefix0("6"))
	fmt.Println(AddPrefix0("9"))
	fmt.Println(AddPrefix0("10"))
	fmt.Println(AddPrefix0("11"))
	fmt.Println(AddPrefix0("12"))
}

func TestString3(t *testing.T) {
	str := "06"
	str = AddPrefix0(str)
	fmt.Println(str)
	fmt.Println(trimPrefix0(str))
	fmt.Println(str)
}

func convertToDecimal1(num int) float64 {
	numStr := strconv.Itoa(num)
	decimalStr := "0." + numStr
	decimal, _ := strconv.ParseFloat(decimalStr, 64)
	return decimal
}

func convertToDecimal2(num int) string {
	return cast.ToString(float64(num) / 100000000000)
}

func Test转换小数(t *testing.T) {
	fmt.Println(convertToDecimal1(1))
	fmt.Println(convertToDecimal1(34))
	fmt.Println(convertToDecimal1(73664))
	fmt.Println(convertToDecimal1(0))

	fmt.Println(convertToDecimal2(1334))
	fmt.Println(convertToDecimal2(34))
	fmt.Println(convertToDecimal2(73664))
	fmt.Println(convertToDecimal2(34433244))
}

func TestTrimSpace(t *testing.T) {
	fmt.Print(strings.TrimSpace("  0 12 "))
	fmt.Print("-")
	fmt.Println(strings.TrimSpace("  1"))
	fmt.Println(strings.TrimSpace("  2"))
	fmt.Println(strings.TrimSpace("  3"))
	fmt.Println(strings.TrimSpace("  4"))
	fmt.Println(strings.TrimSpace("  5"))
	fmt.Println(strings.TrimSpace("  6"))
	fmt.Println(strings.TrimSpace("  7"))
	fmt.Println(strings.TrimSpace("  8"))
	fmt.Println(strings.TrimSpace("  9"))
}

func TestStringx(t *testing.T) {
	x := "[[1]]"
	fmt.Println(string(x[2]))
}

func TestCutUrl(t *testing.T) {
	// 给定的URL
	rawURL := "https://static-inc.xiwang.com/xwx-user-avatar/test/1b25e2f50c7024f2be29be34b6fe4718.png"

	// 解析URL
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	// 获取路径部分
	path := parsedURL.Path
	path = strings.TrimPrefix(parsedURL.Path, "/")

	fmt.Println("提取的路径部分为:", path)
}

func TestAppend(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	fmt.Println(append(a, b...))
}

func TestMapFmt(t *testing.T) {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fmt.Printf("%+v\n", m)
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().UnixMilli())
	fmt.Println(time.Now().UnixMicro())
	fmt.Println(time.Now().UnixNano())
}

type T struct {
	A int
}

func TestSlice1(t *testing.T) {
	var s []*T
	fmt.Println(s)
	s = make([]*T, 0) // 未初始化的切片为nil
	slice(s)
}

func slice(t []*T) {
	fmt.Println(t)
	if t == nil {
		fmt.Println("t is nil")
	} else {
		fmt.Println("t is not nil")
	}
}

func TestStrContains(t *testing.T) {
	fmt.Println(strings.Contains("1,2,3,4,5", "3"))
	fmt.Println(strings.Contains("1,2,3,4,5", "6"))
	fmt.Println(strings.Contains("你好%s", "%s"))
	fmt.Println(strings.Contains("%s还未完成，快去完成吧！", "%s"))
	s1 := make([]int, 0)
	s2 := []int{}
	fmt.Println(s1)
	fmt.Println(s1 == nil)
	fmt.Println(s2)
	fmt.Println(s2 == nil)
}

func Test20250210(t *testing.T) {
	fmt.Println("New Year First Code Day!")
}

var p *int

func foo() (*int, error) {
	var i int = 5
	return &i, nil
}

func bar() {
	// use p
	fmt.Println(*p)
}

func TestPoint(t *testing.T) {
	// 对于使用:=定义的变量，如果新变量与同名已定义的变量不在同一个作用域中，那么 Go 会新定义这个变量。
	p, err := foo() // 此处的p会覆盖全局变量定义的p
	if err != nil {
		fmt.Println(err)
		return
	}
	bar()
	fmt.Println(*p)
}

// 问题代码
func TestGoThread(t *testing.T) {
	x := []int{1, 2, 3, 4, 5}
	var y []int
	for i := range x {
		go func() {
			y = append(y, i)
		}()
	}
	fmt.Println(y)
}

// 代码解决方案
/*
1. 循环变量捕获问题
2. 并发写入的数据竞争 切片不是并发安全的
3. 缺乏同步导致结果不可预测
*/
func TestGoThreadSolution(t *testing.T) {
	x := []int{1, 2, 3, 4, 5}
	var (
		y  []int
		mu sync.Mutex
		wg sync.WaitGroup
	)

	for i := range x {
		z := i // 重新生成变量或者传入值也可以
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			y = append(y, z)
		}()
	}
	wg.Wait()
	fmt.Println(y) // 输出结果可能是乱序的，但包含所有元素
}

func appendStr() func(string) string {
	t := "Hello"
	z := func(x string) string {
		t = t + " " + x
		return t
	}
	t = "wwww" // 闭包引用的是外部变量的地址，这里修改也会影响闭包里面的t的值
	return z
}

func TestPrint(t *testing.T) {
	a := appendStr()
	b := appendStr()
	fmt.Println(a("world"))
	fmt.Println(b("Go"))
	fmt.Println(a("ZTy"))
	fmt.Println(b("!"))
}

func TestInit(t *testing.T) {
	// 对于nil的slice，能够直接append数据。但是不能直接访问数据，否则会panic
	var arr []int
	// nil的map和slice不同，是不能直接赋值的，会panic，需要使用make初始化，但是可以直接访问
	var m map[int]int
	fmt.Println(arr)
	fmt.Println(m)
	// fmt.Println(arr[0])
	fmt.Println(m[0])
}

func create() (fs [2]func()) {
	for i := 0; i < 2; i++ {
		fs[i] = func() {
			fmt.Println(i)
		}
	}
	return
}
func TestClosure2(t *testing.T) {
	fs := create()
	// 闭包引用的是外部变量的地址，所以会打印最后的i值，共用的是同一个i
	for i := 0; i < len(fs); i++ {
		fs[i]()
	}
	s := "【上】讲次2"
	fmt.Println(strings.TrimPrefix(s, "【上】"))
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// 二叉树的最大深度
func TestGetMaxLen(t *testing.T) {
	root := &Node{
		Val: 1,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val: 4,
			},
			Right: &Node{
				Val: 5,
			},
		},
		Right: &Node{
			Val: 3,
		},
	}
	t.Log(getMaxLen(root))
}

func getMaxLen(root *Node) int {
	if root == nil {
		return 0
	}
	left := getMaxLen(root.Left)
	right := getMaxLen(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}

// 二叉树的最大深度路径数据
func TestGetMaxLenPath(t *testing.T) {
	root := &Node{
		Val: 1,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val: 4,
			},
			Right: &Node{
				Val: 5,
			},
		},
		Right: &Node{
			Val: 3,
		},
	}
	t.Log(getMaxLenPath(root))
}

func getMaxLenPath(root *Node) []int {
	if root == nil {
		return []int{}
	}
	left := getMaxLenPath(root.Left)
	right := getMaxLenPath(root.Right)
	if len(left) > len(right) {
		return append([]int{root.Val}, left...)
	}
	return append([]int{root.Val}, right...)
}

func appendInt(x []int) {
	x = append(x, 0)
}

func mod(x []int) {
	x[0] = 99
}

func TestSlice22(t *testing.T) {
	x := make([]int, 1)
	fmt.Println(x)
	appendInt(x) // x = append(x, 0)不会改变原数值的值，可以通过转递指针实现 *[]int
	fmt.Println(x)
	mod(x)
	fmt.Println(x)
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(2)
	fmt.Println(randNum)
}

type Mt struct {
	Name string
	Age  int
}

func TestArray(t *testing.T) {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a[1:])
	// fmt.Println([3]int{1, 2, 3}[1:]) // 不能这么寻址
	m := map[string]Mt{
		"a": {
			Name: "zhangsan",
			Age:  18,
		},
		"b": {
			Name: "lisi",
			Age:  20,
		},
	}
	fmt.Println(m["a"])
	fmt.Println(m["b"].Name) // 可以取值 但是不能赋值：m["b"].Name = ""  因为map的元素是不可寻址的

	// map取值得到的是元素值的拷贝副本，内存中不是同一个数据，不能进行修改
	fmt.Printf("%#v\n", m["c"]) // 返回空对象的零值状态 base.Mt{Name:"", Age:0}
}

// 切片Append坑-使用索引访问切片需要注意
/*
当是切片（slice）时，表达式 s[low : high] 中的 high，最大的取值范围对应着切片的容量（cap），不是单纯的长度（len）。因此调用 fmt.Println(sl[:10]) 时可以输出容量范围内的值，不会出现越界。
相对的 fmt.Println(sl) 因为该切片 len 值为 0，没有指定最大索引值，high 则取 len 值，导致输出结果为空。
*/
func TestAppend1(t *testing.T) {
	sl := make([]int, 0, 10)
	var appenFunc = func(s []int) {
		s = append(s, 10, 20, 30) // 这里的s是一个新的切片，不会影响原切片,底层共用一个数组
		fmt.Println(s)            // [10 20 30]
	}
	fmt.Println(sl) // []
	appenFunc(sl)
	fmt.Println(sl)      // [] 长度为0 类似于[0:0] 取不到元素
	fmt.Println(sl[0:0]) // []
	fmt.Println(sl[:10]) // [10 20 30 0 0 0 0 0 0 0]
}

type Doc struct {
	Name     string
	Children *[]Doc
}

// 输出文件树目录结构
func TestDoc(t *testing.T) {
	doc := Doc{
		Name: "root",
		Children: &[]Doc{
			{
				Name: "child1",
				Children: &[]Doc{
					{
						Name: "1",
					},
					{
						Name: "2",
					},
				},
			},
			{
				Name: "child2",
				Children: &[]Doc{
					{
						Name: "1",
					},
					{
						Name: "2",
					},
				},
			},
		},
	}
	// 输出路径树
	bytes, _ := json.Marshal(doc)
	fmt.Println(string(bytes))
	dfs("", doc)
}

func dfs(str string, doc Doc) {
	path := str
	if path == "" {
		path = doc.Name
	} else {
		path = path + "/" + doc.Name
	}
	if doc.Children == nil || len(*doc.Children) == 0 {
		fmt.Println(path)
		return
	}
	for _, v := range *doc.Children {
		dfs(path, v)
	}
}

func TestChannel(t *testing.T) {
	ch := make(chan int)
	setData(ch)
	for x := range ch {
		fmt.Println(x)
	}
	fmt.Println("done")
}

func setData(ch chan int) {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		x := i
		wg.Add(1)
		go func() {
			ch <- x
			time.Sleep(time.Second)
			wg.Done()
		}()
	}
	wg.Wait()
	close(ch)
}

// test 链接1
// https://www.bilibili.com/video/BV1rP4y1A7j4/
// test 链接2
// https://www.bilibili.com/video/BV1rP4y1A7j4?p=2
// test 链接3
// https://www.bilibili.com/video/BV1rP4y1A7j4?p=3

type TraceID struct {
	High, Low uint64
}

// String returns a string representation of the TraceID ---- 使用此方法可以自定义结构体的字符串输出形式
func (t TraceID) String() string {
	if t.High == 0 {
		return fmt.Sprintf("%x", t.Low) // 输出 16 进制
	}
	return fmt.Sprintf("%x%016x", t.High, t.Low)
}

func GetTraceID() TraceID {
	return TraceID{
		High: 0,
		Low:  7771031616213812862,
	}
}

func TestGenStr(t *testing.T) {
	s := fmt.Sprint(GetTraceID())
	fmt.Println(s)
	x := []string{s}
	fmt.Println(x)

	fmt.Println(GetTraceID())
}
