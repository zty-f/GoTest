package algorithm

import (
	"fmt"
	"sort"
	"testing"
)

/*
双指针/滑动窗口
●LeetCode3 无重复字符的最长子串
●LeetCode11 盛最多水的容器
●LeetCode15 三数之和
●LeetCode16 最接近的三数之和
●LeetCode26 删除排序数组中的重复项
●LeetCode42 接雨水
●LeetCode121 买卖股票的最佳时机
●LeetCode209 长度最小的子数组
*/

// 1、LeetCode3 无重复字符的最长子串
/*
给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度。
示例 1:
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:
输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
*/

func TestLeetCode3(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring1("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring1("bbbbb"))
	fmt.Println(lengthOfLongestSubstring2("pwwkew"))
	fmt.Println(lengthOfLongestSubstring2("aab"))
}

// 方法一： 使用map
func lengthOfLongestSubstring1(s string) int {
	if len(s) == 0 {
		return 0
	}
	cm := make(map[byte]struct{})
	left, right := 0, 0
	res := 0
	for right < len(s) {
		if _, ok := cm[s[right]]; !ok {
			cm[s[right]] = struct{}{}
			right++
		} else {
			if len(cm) > res {
				res = len(cm)
			}
			_, ex := cm[s[right]]
			for ex {
				delete(cm, s[left])
				left++
				_, ex = cm[s[right]]
			}
		}
	}
	if len(cm) > res {
		res = len(cm)
	}
	return res
}

// 方法二：使用数组--推荐
func lengthOfLongestSubstring2(s string) int {
	if len(s) == 0 {
		return 0
	}
	cm := make([]int, 128)
	left := 0
	res := 0
	for right := 0; right < len(s); right++ {
		cm[s[right]]++
		for cm[s[right]] > 1 {
			cm[s[left]]--
			left++
		}
		if right-left+1 > res {
			res = right - left + 1
		}
	}
	return res
}

// 方法三：滑动窗口 每次只需要把left重置到当前重复的地方的下一个位置即可-推荐
func lengthOfLongestSubstring3(s string) int {

	res := 0
	posMap := make(map[byte]int)
	size := len(s)
	l := 0
	for r := 0; r < size; r++ {
		if v, ok := posMap[s[r]]; ok {
			if v+1 > l {
				l = v + 1
			}
		}
		if r-l+1 > res {
			res = r - l + 1
		}
		posMap[s[r]] = r
	}
	return res
}

// 2、LeetCode11 盛最多水的容器
/*
给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
返回容器可以储存的最大水量。
说明：你不能倾斜容器。
输入：[1,8,6,2,5,4,8,3,7]
输出：49
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
示例 2：
输入：height = [1,1]
输出：1
*/

func TestLeetCode11(t *testing.T) {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea2([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

// 方法一：两次循环 O(n^2) 超出时间限制
func maxArea(height []int) int {
	res := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			res = max(res, min(height[i], height[j])*(j-i))
		}
	}
	return res
}

// 方法二：双指针
/*
S(i,j)=min(h[i],h[j])×(j−i)
在每个状态下，无论长板或短板向中间收窄一格，都会导致水槽 底边宽度 −1 变短：
若向内 移动短板 ，水槽的短板 min(h[i],h[j]) 可能变大，因此下个水槽的面积 可能增大 。
若向内 移动长板 ，水槽的短板 min(h[i],h[j]) 不变或变小，因此下个水槽的面积 一定变小 。
因此，初始化双指针分列水槽左右两端，循环每轮将短板向内移动一格，并更新面积最大值，直到两指针相遇时跳出；即可获得最大面积。
*/
func maxArea2(height []int) int {
	left, right, res := 0, len(height)-1, 0
	for left < right {
		res = max(res, min(height[left], height[right])*(right-left))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return res
}

// 3、LeetCode15 三数之和
/*
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。
注意：答案中不可以包含重复的三元组。
示例 1：
输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
解释：
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
注意，输出的顺序和三元组的顺序并不重要。
*/

func TestLeetCode15(t *testing.T) {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4})) // -4 -1 -1 0 1 2
	fmt.Println(threeSum([]int{0, 1, 1}))
}

// 先固定一个数，然后使用双指针，注意重复数字的处理
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	size := len(nums)
	for i := 0; i < size-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i]+nums[i+1]+nums[i+2] > 0 { // 优化一 三数之和大于0，后面的数不可能等于0
			break
		}
		if nums[i]+nums[size-2]+nums[size-1] < 0 { // 优化二 当前数加上最大的两个数小于0，不可能等于0
			continue
		}
		j, k := i+1, size-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum < 0 {
				j++
			} else if sum > 0 {
				k--
			} else {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				for j++; j < k && nums[j] == nums[j-1]; j++ {
				}
				for k--; j < k && nums[k] == nums[k+1]; k-- {
				}
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
