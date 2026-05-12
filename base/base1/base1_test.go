package base1

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 输入： 1->2
// 输出： false
//
// 示例 2：
//
// 输入： 1->2->2->1
// 输出： true
// 判断链表是否为回文链表
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	if head.Next == nil {
		return true
	}
	slow, fast := head, head
	var pre *ListNode
	// 1 2 3 4 4 3 2 1
	// 1 2 3 4 3 2 1
	// 0 0
	// 1 0 1
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		nxt := slow.Next
		slow.Next = pre
		pre = slow
		slow = nxt
	}
	mid := slow
	if fast != nil {
		mid = mid.Next
	}
	// 3 2 1 4 3 2 1
	for mid != nil && pre != nil {
		if mid.Val != pre.Val {
			return false
		}
		// 还原链表
		nxt := pre
		mid, pre = mid.Next, pre.Next
		nxt.Next, slow = slow, nxt
	}
	return true
}

func TestJudgeList(t *testing.T) {
	// head := &ListNode{
	// 	Val: 1,
	// 	Next: &ListNode{
	// 		Val: 2,
	// 		Next: &ListNode{
	// 			Val: 3,
	// 			Next: &ListNode{
	// 				Val: 4,
	// 				Next: &ListNode{
	// 					Val: 4,
	// 					Next: &ListNode{
	// 						Val: 3,
	// 						Next: &ListNode{
	// 							Val: 2,
	// 							Next: &ListNode{
	// 								Val: 1,
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
	// head := &ListNode{
	// 	Val: 0,
	// 	Next: &ListNode{
	// 		Val: 0,
	// 	},
	// }
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 1,
			},
		},
	}
	slow, fast := head, head
	var pre *ListNode
	// 1 2 3 4 4 3 2 1
	// 1 2 3 4 3 2 1
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		nxt := slow.Next
		slow.Next = pre
		pre = slow
		slow = nxt
	}
	for pre != nil {
		fmt.Println(pre.Val)
		pre = pre.Next
	}
}
