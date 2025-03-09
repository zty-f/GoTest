package algorithm

import "testing"

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
		return append(left, root.Val)
	}
	return append(right, root.Val)
}
