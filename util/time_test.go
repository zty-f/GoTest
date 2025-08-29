package util

import (
	"fmt"
	"github.com/spf13/cast"
	"strings"
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
	(*m)[0] = 1 // 为空的，不能直接赋值
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

func TestGetWeekStartAndEnd1(t *testing.T) {
	start, end := GetWeekStartAndEnd1()
	fmt.Println(start, end)
	fmt.Println(time.Now().Weekday())
}

func TestGetWeekStartAndEnd2(t *testing.T) {
	start, end := GetWeekStartAndEnd()
	fmt.Println(start, end)
	fmt.Println(time.Now().Weekday())
}

func TestGetWeekStartAndEnd3(t *testing.T) {
	start, end := GetStartAndEndTimeOfCurrentWeek()
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

func TestStringCut(t *testing.T) {
	str := "【完成条件】\n1.单讲出勤：单节课直播观看时长大于0即视为完成出勤。注意，每节课出勤仅计算一次。若同一节课中，主讲或伴学有任意一方完成出勤行为，即算该讲次出勤。单讲按时出勤可获得10成长值。\n2.全部出勤：当同一学科课程的全部讲次均完成出勤，即每节课的直播观看时长均大于0时，即视为完成全部出勤。全部按时出勤可获得40成长值。\n\n【课程类型】\n可完成该任务的课程类型包括系统直播课（含伴学）、体验课及特训班。\n\n【限时奖励】\n课程当天完成出勤，将额外+5成长值。"
	// 截断【限时奖励】之后的字符串
	str = str[0:strings.Index(str, "【限时奖励】")]
	fmt.Println(str)
}

func TestTimestampToDate(t *testing.T) {
	fmt.Println(TimestampToDate(0 / 1000))
}

func TestGetMonthStartAndEndTime(t *testing.T) {
	type testCase struct {
		ct int64
	}
	cases := []testCase{
		{
			ct: 1642730400, // 2022-01-21 10:00:00
		},
		{
			ct: 1740755514, // 2025-02-28 23:11:54
		},
		{
			ct: 1745305916, // 2025-04-22 15:11:56
		},
		{
			ct: 1735369914, // 2024-12-28 15:11:54
		},
		{
			ct: 1709190714, // 2024-02-29 15:11:54
		},
	}
	for _, c := range cases {
		fmt.Println(GetMonthStartAndEndTime(time.Unix(c.ct, 0)))
	}
}

func TestGetNextDay8AM(t *testing.T) {
	date := "2025-05-10 23:59:59"
	// 将字符串转换为时间对象
	now, _ := time.Parse("2006-01-02 15:04:05", date)
	nextDay8AM := GetNextDay8AM(now)
	fmt.Printf("当前时间: %s\n", now.Format(time.RFC3339))
	fmt.Printf("第二天 8 点的时间: %s\n", nextDay8AM.Format(time.RFC3339))
}

func TestSecondsToMinutes(t *testing.T) {
	fmt.Println(SecondsToMinutes(60))
	fmt.Println(SecondsToMinutes(245))
	fmt.Println(SecondsToMinutes(234))
	fmt.Println(SecondsToMinutes(3334))
	fmt.Println(SecondsToMinutes(54673))
	fmt.Println(fmt.Sprintf("%.2f", SecondsToMinutes(54673)))
	fmt.Println(fmt.Sprintf("%.2f", 2.578)) // 会四舍五入
	fmt.Println(fmt.Sprintf("%.2f", 2.573))
}

func TestGetDayOfWeek(t *testing.T) {
	fmt.Println(GetDayOfWeek(time.Now()))
	fmt.Println(GetDayOfWeek(time.Date(2025, 8, 25, 0, 0, 0, 0, time.Local)))
	fmt.Println(GetDayOfWeek(time.Date(2025, 8, 26, 0, 0, 0, 0, time.Local)))
	fmt.Println(GetDayOfWeek(time.Date(2025, 8, 27, 0, 0, 0, 0, time.Local)))
	fmt.Println(GetDayOfWeek(time.Date(2025, 8, 28, 0, 0, 0, 0, time.Local)))
	fmt.Println(GetDayOfWeek(time.Date(2025, 8, 29, 0, 0, 0, 0, time.Local)))
	fmt.Println(GetDayOfWeek(time.Date(2025, 8, 30, 0, 0, 0, 0, time.Local)))
	fmt.Println(GetDayOfWeek(time.Date(2025, 8, 31, 0, 0, 0, 0, time.Local)))
}

func TestIsValidTimeFormat(t *testing.T) {
	fmt.Println(IsValidTimeFormat("18:00:00")) // true
	fmt.Println(IsValidTimeFormat("-1:00:00")) // true
	fmt.Println(IsValidTimeFormat("18:76:00")) // true
	fmt.Println(IsValidTimeFormat("25:00:00")) // false
	fmt.Println(IsValidTimeFormat("24:00:00")) // false
	fmt.Println(IsValidTimeFormat("18:60:00")) // false
	fmt.Println(IsValidTimeFormat("18:00"))    // false
}
