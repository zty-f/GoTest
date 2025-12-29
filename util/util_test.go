package util

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func Test1(t *testing.T) {
	// num := 0007
	// number := strconv.FormatInt(int64(num), 2)
	// fmt.Println(number)
	number := "42"
	res, err := strconv.ParseUint(number, 2, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
	// ParseInt实质上先对于字符串数字前面的+-符号进行处理，然后再调用ParseUint方法
	res1, err := strconv.ParseInt(number, 2, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res1)
}

func TestGetMonthStartTime(t *testing.T) {
	dateStr := "2024-01"
	time, _ := GetMonthStartTime(dateStr)
	fmt.Println(time)
	time, _ = GetMonthEndTime(dateStr)
	fmt.Println(time)
}

// 生成随机字符串
var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}
func GetRandomString1(l int) string {
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	var result []byte
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func GetRandomString2(l int) string {
	bytes := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	var result = make([]byte, l)
	for i := range result {
		result[i] = bytes[r.Intn(len(bytes))]
	}
	return string(result)
}

func GetRandomString3(n int) string {
	randBytes := make([]byte, n/2)
	rand.Read(randBytes)
	return fmt.Sprintf("%x", randBytes)
}

func TestGetRandomString(t *testing.T) {
	fmt.Println(GetRandomString1(10))
	fmt.Println(GetRandomString2(10))
	fmt.Println(GetRandomString3(10))
}

func BenchmarkGetRandomString1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetRandomString1(10)
	}
}

func BenchmarkGetRandomString2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetRandomString2(10)
	}
}

func BenchmarkGetRandomString3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetRandomString3(10)
	}
}

func TestContains(t *testing.T) {
	fmt.Println(Contains([]int{1, 2, 3, 4}, 3))
	fmt.Println(Contains([]int{1, 2, 3, 4}, 5))
	fmt.Println(Contains([]string{"a", "b", "c"}, "b"))
	fmt.Println(Contains([]string{"a", "b", "c"}, "d"))
}

func TestNumToChinese(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, num := range nums {
		chineseNum, err := NumToChinese(num)
		if err != nil {
			fmt.Printf("Error converting %d to Chinese: %v\n", num, err)
		} else {
			fmt.Printf("%d in Chinese is: %s\n", num, chineseNum)
		}
	}
}

func TestChineseToNum(t *testing.T) {
	chineseNums := []string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九", "十"}
	for _, chineseNum := range chineseNums {
		num, err := ChineseToNum(chineseNum)
		if err != nil {
			fmt.Printf("Error converting %s to number: %v\n", chineseNum, err)
		} else {
			fmt.Printf("%s in number is: %d\n", chineseNum, num)
		}
	}
}

func TestExtractChineseNumber(t *testing.T) {
	// 示例描述
	descriptions := []string{
		"这是第七周的任务",
		"今天是第二周的开始",
		"第十周的会议安排在下周",
		"没有中文数字的描述",
		"没有第五周中文数字的第1周描述第七周",
	}

	for _, desc := range descriptions {
		chineseNum, err := ExtractChineseNumber(desc)
		if err != nil {
			fmt.Printf("Error extracting Chinese number from '%s': %v\n", desc, err)
		} else {
			fmt.Printf("Extracted Chinese number from '%s': %s\n", desc, chineseNum)
		}
	}
}

func TestExtractAndIncrementChineseWeek(t *testing.T) {
	// 示例描述
	descriptions := []string{
		"这是第七周的任务",
		"今天是第二周的开始",
		"第十周的会议安排在下周",
		"没有中文数字的描述",
		"没有第五周中文数字的第1周描述第七周",
	}

	for _, desc := range descriptions {
		newDesc, err := ExtractAndIncrementChineseWeek(desc)
		if err != nil {
			fmt.Printf("Error incrementing Chinese week in '%s': %v\n", desc, err)
		} else {
			fmt.Printf("New description after incrementing week in '%s': %s\n", desc, newDesc)
		}
	}
}

type Conf struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func TestRandomArray(t *testing.T) {
	fmt.Println(RandomArray([]int{1, 2, 3, 4, 5}))
	fmt.Println(RandomArray([]int{1, 2, 3, 4, 5}))
	fmt.Println(RandomArray([]int{1, 2, 3, 4, 5}))
	fmt.Println(RandomArray([]string{"a", "b", "c"}))
	fmt.Println(RandomArray([]string{"a", "b", "c"}))
	fmt.Println(RandomArray([]string{"a", "b", "c"}))
	fmt.Println(RandomArray([]Conf{{Id: 1, Name: "a"}, {Id: 2, Name: "b"}, {Id: 3, Name: "c"}}))
	fmt.Println(RandomArray([]Conf{{Id: 1, Name: "a"}, {Id: 2, Name: "b"}, {Id: 3, Name: "c"}}))
	fmt.Println(RandomArray([]Conf{{Id: 1, Name: "a"}, {Id: 2, Name: "b"}, {Id: 3, Name: "c"}}))
}
