package algorithm

import (
	"fmt"
	"math"
	"testing"
)

// 1. [3]无重复字符的最长子串
func TestLengthOfLongestSubstring(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

func lengthOfLongestSubstring(s string) int {
	/*简单方法，双指针
	l := len(s)
	res := 0
	set := make(map[rune]int)
	i, j := 0, 0
	tmp := 0
	for i <= j && i < l && j < l {
		if _, ok := set[rune(s[j])]; !ok {
			tmp++
			set[rune(s[j])] = 1
			j++
		} else {
			i++
			j = i
			set = make(map[rune]int)
			res = int(math.Max(float64(res), float64(tmp)))
			tmp = 0
		}
	}
	res = int(math.Max(float64(res), float64(tmp)))
	return res*/

	// 滑动窗口 每次只需要把left重置到当前重复的地方的下一个位置即可
	l := len(s)
	res := 0
	set := make(map[rune]int)
	left := 0
	for i := 0; i < l; i++ {
		if v, ok := set[rune(s[i])]; ok {
			left = int(math.Max(float64(left), float64(v+1)))
		}
		res = int(math.Max(float64(res), float64(i-left+1)))
		set[rune(s[i])] = i
	}
	return res
}

// 2. [5]最长回文子串
func TestLongestPalindrome(t *testing.T) {
	fmt.Println(longestPalindrome("tattarrattat"))
}

func longestPalindrome(s string) string {
	size := 1
	l := len(s)
	if l < 2 {
		return s
	}
	res := s[0:1]
	for i := 0; i < l; i++ {
		m := i
		n := i + 1
		for m >= 0 && n < l {
			if s[m] == s[n] {
				if n-m+1 > size {
					res = s[m : n+1]
					size = n - m + 1
				}
				m--
				n++
			} else {
				break
			}
		}
		m = i - 1
		n = i + 1
		for m >= 0 && n < l {
			if s[m] == s[n] {
				if n-m+1 > size {
					res = s[m : n+1]
					size = n - m + 1
				}
				m--
				n++
			} else {
				break
			}
		}
	}
	return res
}

// 3. [2673]使二叉树所有路径值相等的最小代价
/*
给你一个整数 n 表示一棵 满二叉树 里面节点的数目，节点编号从 1 到 n 。根节点编号为 1 ，树中每个非叶子节点 i 都有两个孩子，分别是左孩子 2 * i 和右孩子 2 * i + 1 。
树中每个节点都有一个值，用下标从 0 开始、长度为 n 的整数数组 cost 表示，其中 cost[i] 是第 i + 1 个节点的值。每次操作，你可以将树中 任意 节点的值 增加 1 。你可以执行操作 任意 次。
你的目标是让根到每一个 叶子结点 的路径值相等。请你返回 最少 需要执行增加操作多少次。
注意：
满二叉树 指的是一棵树，它满足树中除了叶子节点外每个节点都恰好有 2 个子节点，且所有叶子节点距离根节点距离相同。
路径值 指的是路径上所有节点的值之和。
       1-1
  2-5	    3-2
4-2	5-3	 6-3	7-1
*/
func TestMinIncrements(t *testing.T) {
	fmt.Println(minIncrements(7, []int{1, 5, 2, 2, 3, 3, 1}))
}

func minIncrements(n int, cost []int) int {
	res := 0
	for i := n - 2; i >= 0; i -= 2 {
		res += abs(cost[i] - cost[i+1])
		cost[i/2] += max(cost[i], cost[i+1])
	}
	return res
}

// 4. [27]移除元素
/*
给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素。元素的顺序可能发生改变。然后返回 nums 中与 val 不同的元素的数量。

假设 nums 中不等于 val 的元素数量为 k，要通过此题，您需要执行以下操作：

更改 nums 数组，使 nums 的前 k 个元素包含不等于 val 的元素。nums 的其余元素和 nums 的大小并不重要。
返回 k。
*/
func TestRemoveElement(t *testing.T) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement1(nums, 2))
	fmt.Println(nums)
	nums = []int{3, 2, 2, 3}
	fmt.Println(removeElement2(nums, 3))
	fmt.Println(nums)
	nums = []int{4}
	fmt.Println(removeElement1(nums, 4))
	fmt.Println(nums)
}

// 这样的方法两个指针在最坏的情况下合起来只遍历了数组一次。 与方法二不同的是，方法一避免了需要保留的元素的重复赋值操作。
func removeElement1(nums []int, val int) int {
	i, j := 0, len(nums)
	for i < j {
		if nums[i] == val {
			nums[i], nums[j-1] = nums[j-1], nums[i]
			j--
		} else {
			i++
		}
	}
	return i
}

// 这样的算法在最坏情况下（输入数组中没有元素等于 val），左右指针各遍历了数组一次。
func removeElement2(nums []int, val int) int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

// LCR 192. 把字符串转换成整数 (atoi)
/*
函数 myAtoi(string s) 的算法如下：
1、读入字符串并丢弃无用的前导空格
2、检查下一个字符（假设还未到字符末尾）为正还是负号，读取该字符（如果有）。 确定最终结果是负数还是正数。 如果两者都不存在，则假定结果为正。
3、读入下一个字符，直到到达下一个非数字字符或到达输入的结尾。字符串的其余部分将被忽略。
4、将前面步骤读入的这些数字转换为整数（即，"123" -> 123， "0032" -> 32）。如果没有读入数字，则整数为 0 。必要时更改符号（从步骤 2 开始）。
5、如果整数数超过 32 位有符号整数范围 [−231,  231 − 1] ，需要截断这个整数，使其保持在这个范围内。具体来说，小于 −231 的整数应该被固定为 −231 ，大于 231 − 1 的整数应该被固定为 231 − 1 。
6、返回整数作为最终结果。
*/
func TestMyAtoi(t *testing.T) {
	fmt.Println(myAtoi("42"))
	fmt.Println(myAtoi("   -42"))
	fmt.Println(myAtoi("4193 with words"))
}

func myAtoi(s string) int {
	// 1. 去除空格
	index := 0
	for index < len(s) && s[index] == ' ' {
		index++
	}
	// 2. 判断正负号
	sign := 1 // 默认是正数
	if index < len(s) && (s[index] == '-' || s[index] == '+') {
		if s[index] == '-' {
			sign = -1
		}
		index++
	}
	res := 0
	for index < len(s) && s[index] >= '0' && s[index] <= '9' {
		res = res*10 + int(s[index]-'0')
		index++
		if res*sign < math.MinInt32 {
			return math.MinInt32
		}
		if res*sign > math.MaxInt32 {
			return math.MaxInt32
		}
	}
	return res * sign
}
