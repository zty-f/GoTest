package main

import (
	"fmt"
)

type A interface {
	ShowA() int
}

type B interface {
	ShowB() int
}

type Work struct {
	i int
}

func (w Work) ShowA() int {
	return w.i + 10
}

func (w Work) ShowB() int {
	return w.i + 20
}

func main() {
	c := Work{3}
	var a A = c
	var b B = c
	//fmt.Println(a.ShowB())  a的静态类型是A，所以ShowB()不是A的方法，不能调用
	//fmt.Println(b.ShowA())  b的静态类型是B，所以ShowA()不是B的方法，不能调用
	fmt.Println(a.ShowA())
	fmt.Println(b.ShowB())

	// 32 位机器
	var x int32 = 32.0
	// 【正确】字面量是无类型（untyped）的，32.0 是无类型的浮点数字面量，因此它可以赋值给任意数字相关类型变量（或常量）
	fmt.Println(x)

	//var y int = x
	//  【错误】Go 语言不会做隐式类型转换，int 和 int32 是不同的类型，因此上题中 2）编译不通过。

	var z rune = x

	fmt.Println(z)
	// 【正确】rune 是 int32 的别名，因此题目中 3）也能编译通过

	str := "hello，你好！"
	fmt.Println(string(str[0]))
	fmt.Println(string(str[1]))
	forbianli(str)
	forrangebianli(str)

	var s1 []int
	var s2 = []int{}
	if s1 == nil {
		fmt.Println("s1 yes nil")
	} else {
		fmt.Println("s1 no nil")
	}

	if s2 == nil {
		fmt.Println("s2 yes nil")
	} else {
		fmt.Println("s2 no nil")
	}
}

func forbianli(s string) {
	str2 := []rune(s)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("str2[%d]= %c\n", i, str2[i])
	}
}

func forrangebianli(s string) {
	for index, val := range s {
		fmt.Printf("index=%d val=%c\n", index, val)
	}
}

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	var res = &ListNode{
		Next: head,
	}
	head = res
	for head.Next != nil && head.Next.Next != nil {
		pre := head
		after := head.Next.Next.Next
		head = head.Next
		pre.Next = head.Next
		pre.Next.Next = head
		head = head.Next
		head.Next = after
	}
	return res.Next
}

func twoSum(nums []int, target int) []int {
	var set map[int]int
	for i := 0; i < len(nums); i++ {
		if v, ok := set[target-nums[i]]; ok {
			return []int{v, i}
		}
		set[nums[i]] = i
	}
	return []int{}
}
