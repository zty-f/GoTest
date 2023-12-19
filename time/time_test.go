package util

import (
	"fmt"
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
		getBinaryBits2(888)
	}
}

func BenchmarkGetBinaryBits2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getBinaryBits2(888)
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
