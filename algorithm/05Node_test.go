package algorithm

import "testing"

/*
树的遍历
●LeetCode94 二叉树的中序遍历
●LeetCode102 二叉树的层次遍历
●LeetCode110 平衡二叉树
●LeetCode144 二叉树的前序遍历
●LeetCode145 二叉树的后序遍历
二叉搜索树
●LeetCode98 验证二叉搜索树
●LeetCode450 删除二叉搜索树中的节点
●LeetCode701 二叉搜索树中的插入操作
递归
●LeetCode21 合并两个有序链表
●LeetCode101 对称二叉树
●LeetCode104 二叉树的最大深度
●LeetCode226 翻转二叉树
●LeetCode236 二叉树的最近公共祖先
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 1、LeetCode94 二叉树的中序遍历\LeetCode144 二叉树的前序遍历\LeetCode145 二叉树的后序遍历
func TestLeetCode94(t *testing.T) {
	root := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	t.Log(inorderTraversal(root))
}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	dfs1(root, &res) // append修改后原值不会变，需要传地址
	return res
}

func dfs1(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	dfs1(root.Left, res)
	*res = append(*res, root.Val)
	dfs1(root.Right, res)
}

// 中序
func inorderTraversal1(root *TreeNode) []int {
	res := make([]int, 0)
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		res = append(res, root.Val)
		dfs(root.Right)
	}
	dfs(root)
	return res
}

// 前序
func inorderTraversal2(root *TreeNode) []int {
	res := make([]int, 0)
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		res = append(res, root.Val)
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return res
}

// 后序
func inorderTraversal3(root *TreeNode) []int {
	res := make([]int, 0)
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		dfs(root.Right)
		res = append(res, root.Val)
	}
	dfs(root)
	return res
}

// 2、LeetCode102 二叉树的层次遍历
func TestLeetCode102(t *testing.T) {
	root := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	t.Log(levelOrder(root))
}

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		tList := make([]int, 0)
		tQueue := make([]*TreeNode, 0)
		for i := 0; i < len(queue); i++ {
			t := queue[i]
			tList = append(tList, t.Val)
			if t.Left != nil {
				tQueue = append(tQueue, t.Left)
			}
			if t.Right != nil {
				tQueue = append(tQueue, t.Right)
			}
		}
		queue = tQueue
		res = append(res, tList)
	}
	return res
}
