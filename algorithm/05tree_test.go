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
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 1、LeetCode94、LeetCode144、LeetCode145
func TestLeetCode94(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
	}
	/*
	    1
	   / \
	  2   2
	     / \
	    3   4
	*/
	// 中序
	t.Log(inorderTraversal1(root))
	t.Log(inorderTraversal2(root))
	// 前序
	t.Log(preorderTraversal(root))
	// 后序
	t.Log(postorderTraversal(root))
}

// 中序
func inorderTraversal1(root *TreeNode) []int {
	res := make([]int, 0)
	dfs(root, &res) // append修改后原值不会变，需要传地址
	return res
}

func dfs(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	dfs(root.Left, res)
	*res = append(*res, root.Val)
	dfs(root.Right, res)
}

// 中序
func inorderTraversal2(root *TreeNode) []int {
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
func preorderTraversal(root *TreeNode) []int {
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
func postorderTraversal(root *TreeNode) []int {
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

// 3、LeetCode110 平衡二叉树
// 平衡二叉树：任意节点的左右子树高度差不超过1
func TestLeetCode110(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
	}
	/*
	    1
	   / \
	  2   2
	     / \
	    3   4
	*/
	t.Log(isBalanced(root))
}

func isBalanced(root *TreeNode) bool {
	return height(root) >= 0
}

// 自底向上 不断更新节点的高度
func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lh := height(root.Left)  // 获取左子树高度
	rh := height(root.Right) // 获取右子树高度
	// 左子树不是平衡二叉树，返回-1
	// 右子树不是平衡二叉树，返回-1
	// 左右子树高度差大于1，返回-1
	if lh == -1 || rh == -1 || abs(lh-rh) > 1 {
		return -1
	}
	return max(lh, rh) + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
