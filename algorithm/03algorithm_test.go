package algorithm

import (
	"fmt"
	"testing"
)

/*
数组操作专栏
1、LeetCode54 螺旋矩阵
2、LeetCode76 最小覆盖子串
3、LeetCode75 颜色分类
4、LeetCode73 矩阵置零
5、LeetCode384 打乱数组
6、LeetCode581 最短无序连续子数组
7、LeetCode945 使数组唯一的最小增量
*/

// 1、LeetCode54 螺旋矩阵
/*
给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]
1 2 3 4 4
4 5 6 5 4
7 8 9 7 5
6 7 8 9 2
*/

func Test螺旋矩阵(t *testing.T) {
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	fmt.Println(spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))
}

func spiralOrder(matrix [][]int) []int {
	i, j := 0, 0
	m, n := len(matrix), len(matrix[0])
	res := make([]int, 0)
	for {
		// 向右
		for j < n && matrix[i][j] != 111 {
			res = append(res, matrix[i][j])
			matrix[i][j] = 111
			if len(res) == m*n {
				return res
			}
			j++
		}
		j--
		i++
		// 向下
		for i < m && matrix[i][j] != 111 {
			res = append(res, matrix[i][j])
			matrix[i][j] = 111
			if len(res) == m*n {
				return res
			}
			i++
		}
		i--
		j--
		// 向左
		for j >= 0 && matrix[i][j] != 111 {
			res = append(res, matrix[i][j])
			matrix[i][j] = 111
			if len(res) == m*n {
				return res
			}
			j--
		}
		j++
		i--
		// 向上
		for i >= 0 && matrix[i][j] != 111 {
			res = append(res, matrix[i][j])
			matrix[i][j] = 111
			if len(res) == m*n {
				return res
			}
			i--
		}
		i++
		j++
	}
}

// 2、LeetCode76 最小覆盖子串
/*
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
注意：
对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
如果 s 中存在这样的子串，我们保证它是唯一的答案。
示例 1：
输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"
解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。
*/

func Test最小覆盖子串(t *testing.T) {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindow("a", "a"))
	fmt.Println(minWindow("a", "aa"))
	fmt.Println(minWindow("abbcbsdaa", "aa"))
	fmt.Println(minWindow("a", "b"))
}

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	cnt := len(t)
	need := make(map[byte]int)
	for i := 0; i < cnt; i++ {
		need[t[i]]++
	}
	left, right := 0, 0
	start, end := -1, 0

	for right < len(s) {
		need[s[right]]--
		if need[s[right]] >= 0 {
			cnt--
		}
		if cnt == 0 {
			for left < right && need[s[left]] < 0 {
				need[s[left]]++
				left++
			}
			if right-left < end-start || end == 0 {
				start, end = left, right
			}
			need[s[left]]++
			left++
			cnt++
		}
		right++
	}
	if start == -1 {
		return ""
	}
	return s[start : end+1]
}

// 3、LeetCode75 颜色分类
/*
给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums，（原地）对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
必须在不使用库内置的 sort 函数的情况下解决这个问题。
示例 1：
输入：nums = [2,0,2,1,1,0]
输出：[0,0,1,1,2,2]
示例 2：
输入：nums = [2,0,1]
输出：[0,1,2]
*/

func Test颜色分类(t *testing.T) {
	x1 := []int{2, 0, 2, 1, 1, 0}
	sortColors1(x1)
	fmt.Println(x1)
	x2 := []int{2, 0, 1}
	sortColors2(x2)
	fmt.Println(x2)
	x3 := []int{1, 0, 2}
	sortColors2(x3)
	fmt.Println(x3)
}

// 方法一，双指针
func sortColors1(nums []int) {
	left := 0
	right := len(nums) - 1
	index := 0
	for index <= right {
		if nums[index] == 2 {
			swap(nums, index, right)
			right--
		} else if nums[index] == 0 {
			swap(nums, left, index)
			left++
			index++
		} else {
			index++
		}
	}
}

// 方法二 两次遍历
func sortColors2(nums []int) {
	x := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			swap(nums, i, x)
			x++
		}
	}
	for i := x; i < len(nums); i++ {
		if nums[i] == 1 {
			swap(nums, i, x)
			x++
		}
	}
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

// 4、LeetCode73 矩阵置零
/*
给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。
输入：matrix = [[1,1,1],[1,0,1],[1,1,1]]
输出：[[1,0,1],[0,0,0],[1,0,1]]
1 1 1
1 0 1
1 1 1
*/
func Test矩阵置零(t *testing.T) {
	x1 := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes1(x1)
	fmt.Println(x1)
	x2 := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes2(x2)
	fmt.Println(x2)
}

// 方法一：两次循环，第一次循环记录需要置零的行和列，第二次循环置零 空间复杂度O(m+n)
func setZeroes1(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	row, col := make([]bool, m), make([]bool, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if row[i] || col[j] {
				matrix[i][j] = 0
			}
		}
	}
}

// 方法二：
func setZeroes2(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	row0, col0 := false, false
	// 遍历第一行，判断是否有0
	for _, v := range matrix[0] {
		if v == 0 {
			row0 = true
			break
		}
	}
	// 遍历第一列，判断是否有0
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			col0 = true
			break
		}
	}
	// 遍历除第一行和第一列的其他行列，如果有0，则将对应的第一行和第一列置为0
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	// 遍历除第一行和第一列的其他行列，如果对应的第一行和第一列为0，则将当前行列置为0
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	// 如果第一行有0，则将第一行置为0
	if row0 {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}
	// 如果第一列有0，则将第一列置为0
	if col0 {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}
