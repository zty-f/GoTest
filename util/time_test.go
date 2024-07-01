package util

import (
	"fmt"
	"github.com/spf13/cast"
	"test/constant"
	"testing"
	"time"
)

func TestSub(t *testing.T) {
	now := time.Now()
	todayStartTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	tn := time.Now()
	fmt.Println(tn.Sub(todayStartTime))
	fmt.Println(tn.Sub(todayStartTime) < 10*time.Minute)
	fmt.Println(todayStartTime.Sub(tn))
	fmt.Println(todayStartTime.Sub(tn) < 10*time.Minute)
}

func TestTransFer(t *testing.T) {
	arr := []int{2, 3, 5, 8}
	x := getNumberWithBitsSet1(arr)
	fmt.Println(x)
	y := getNumberWithBitsSet2(arr...)
	fmt.Println(y)
}

func TestTransFer2(t *testing.T) {
	tmp := 356
	x, _ := getBinaryBits1(tmp)
	fmt.Println(x)
	y := getBinaryBits2(tmp)
	fmt.Println(y)
}

func BenchmarkGetBinaryBits1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getBinaryBits1(888)
	}
}

func BenchmarkGetBinaryBits2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getBinaryBits2(888)
	}
}

func BenchmarkGetNumberWithBitsSet1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getNumberWithBitsSet1([]int{1, 2, 3, 4, 5, 6, 7, 8})
	}
}

func BenchmarkGetNumberWithBitsSet2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getNumberWithBitsSet2([]int{1, 2, 3, 4, 5, 6, 7, 8}...)
	}
}

func TestNew(t *testing.T) {
	list := new([]int)
	fmt.Printf("%T\n", list)
	fmt.Println(list)
	// 会扩容
	*list = append(*list, 1)
	fmt.Println(list)
	fmt.Println("------------")
	m := new(map[int]int)
	fmt.Printf("%T\n", m)
	fmt.Println(m)
	(*m)[0] = 1 //为空的，不能直接赋值
	fmt.Println(m)
}

func TestMake(t *testing.T) {
	list := make([]int, 0)
	list = append(list, 1)
	fmt.Println(list)

	a := 1
	fmt.Printf("%T\n", &a)
	b := new(int)
	fmt.Printf("%T\n", b)
}

func deferFuncReturn1() int {
	i := 1
	defer func() {
		i++
		fmt.Println(i)
	}()
	i = 2
	return i
}

func deferFuncReturn2() (j int) {
	i := 1
	defer func() {
		j++
		fmt.Println(i)
		fmt.Println(j)
	}()

	return i
}

func TestDefer(t *testing.T) {
	fmt.Println(deferFuncReturn1())
	fmt.Println("-------------")
	fmt.Println(deferFuncReturn2())
}

func f(x int, y int) int {
	fmt.Println("x:", x)
	return x
}

// defer里面有嵌套函数会先执行
func TestDefer2(t *testing.T) {
	defer f(1, f(3, 0))
	defer f(2, f(4, 0))
	time.Sleep(2 * time.Second)
}

func defer_fun1(x int) (res int) {
	res = x
	defer func() {
		res += 3
	}()
	return res
}

func defer_fun2(x int) int {
	res := x
	defer func() {
		res += 3
	}()
	return res
}

func defer_fun3(x int) (res int) {
	defer func() {
		res += x
	}()
	return 2
}

func defer_fun4() (res int) {
	t := 1
	defer func(x int) {
		fmt.Println("x:", x)
		fmt.Println("res:", res)
	}(t)
	t = 2
	return 4
}

func TestDefer3(t *testing.T) {
	fmt.Println(defer_fun1(1)) // 4
	fmt.Println(defer_fun2(1)) // 1
	fmt.Println(defer_fun3(1)) // 3
	fmt.Println(defer_fun4())  // 1 4 4
}

func TestGetWeekStartAndEnd(t *testing.T) {
	start, end := GetWeekStartAndEnd()
	fmt.Println(start, end)
	fmt.Println(time.Now().Weekday())
}

func TestGetMonday(t *testing.T) {
	monday := GetMonday("0102")
	fmt.Println(monday)
}

// Truncate方法只支持标准时区
func TestTimeTruncate1(t *testing.T) {
	now := time.Now().UTC()
	fmt.Println(now)
	fmt.Println(now.Truncate(24 * time.Hour))
}

func TestTimeTruncate2(t *testing.T) {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Truncate(24 * time.Hour))
}

func TestTimeDuration(t *testing.T) {
	start, _ := StrToTime1("2024-07-01")
	end, _ := StrToTime1("2024-07-08")
	duration := GetTwoDateIntervalDuration(start, end)
	fmt.Println(duration)
	if duration > constant.SevenDay {
		fmt.Println(">大于")
	}
	if duration < constant.SevenDay {
		fmt.Println("<小于")
	}
	if duration == constant.SevenDay {
		fmt.Println("=等于")
	}
}

func TestNOW(t *testing.T) {
	fmt.Println(time.Now().Year())
	fmt.Println(cast.ToString(cast.ToInt(time.Now().Month())))
}

func TestTimeDuration1(t *testing.T) {
	start, _ := StrToTime("2024-07-01 00:00:00")
	fmt.Println(start)
	start = start.Add(constant.SevenDay)
	fmt.Println(start)
}
