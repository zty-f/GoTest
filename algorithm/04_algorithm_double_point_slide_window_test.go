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

// 4、LeetCode16 最接近的三数之和
/*
给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。
返回这三个数的和。假定每组输入只存在恰好一个解。
示例 1：
输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2)。
示例 2：
输入：nums = [0,0,0], target = 1
输出：0
解释：与 target 最接近的和是 0（0 + 0 + 0 = 0）。
*/
func TestLeetCode16(t *testing.T) {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
	fmt.Println(threeSumClosest([]int{0, 0, 0}, 1))
}

// -4 -1 1 2
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	res := 1 << 32
	for i := 0; i < n-2; i++ {
		x := nums[i]
		if i > 0 && x == nums[i-1] {
			continue // 优化三
		}

		// 优化一
		s := x + nums[i+1] + nums[i+2]
		if s > target { // 后面无论怎么选，选出的三个数的和不会比 s 还小
			if abs(s-target) < abs(res-target) {
				res = s
			}
			break
		}
		// 优化二
		s = x + nums[n-2] + nums[n-1]
		if s < target { // x 加上后面任意两个数都不超过 s，所以下面的双指针就不需要跑了
			if abs(s-target) < abs(res-target) {
				res = s
			}
			continue
		}
		j, k := i+1, n-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if abs(sum-target) < abs(res-target) {
				res = sum
			}
			if sum > target {
				k--
			} else if sum < target {
				j++
			} else {
				return target
			}
		}
	}
	return res
}

// 5、LeetCode26 删除排序数组中的重复项
/*
给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，
返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。
考虑 nums 的唯一元素的数量为 k ，你需要做以下事情确保你的题解可以被通过：
更改数组 nums ，使 nums 的前 k 个元素包含唯一元素，并按照它们最初在 nums 中出现的顺序排列。nums 的其余元素与 nums 的大小不重要。
返回 k 。
输入：nums = [1,1,2]
输出：2, nums = [1,2,_]
解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素。
*/
func TestLeetCode26(t *testing.T) {
	nums := []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates1(nums))
	fmt.Println(nums)
	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates2(nums))
	fmt.Println(nums)
}

// 方法一：使用额外空间 + 双指针
func removeDuplicates(nums []int) int {
	set := make(map[int]struct{})
	i := 0
	for j := 0; j < len(nums); j++ {
		if _, ok := set[nums[j]]; !ok {
			nums[i] = nums[j]
			set[nums[j]] = struct{}{}
			i++
		}
	}
	return i
}

// 方法二：原地使用双指针，不额外空间！！！推荐-比较好理解
func removeDuplicates1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 1
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[j-1] {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

// 方法三：原地使用双指针，不额外空间！！！推荐 和上面的略有不同 -比较不好理解
func removeDuplicates2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

// 6、LeetCode42 接雨水
/*
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
示例 1：
输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
示例 2：
输入：height = [4,2,0,3,2,5]
输出：9
*/
func TestLeetCode42(t *testing.T) {
	fmt.Println(trap1([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Println(trap1([]int{4, 2, 0, 3, 2, 5}))
	fmt.Println(trap2([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Println(trap2([]int{4, 2, 0, 3, 2, 5}))
}

// 方法一：双向获取最大值，第i个位置能存储的雨水量 = min(leftMax[i], rightMax[i]) - height[i]
func trap1(height []int) int {
	afterMax := make([]int, len(height))
	n := len(height)
	afterMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		afterMax[i] = max(height[i], afterMax[i+1])
	}
	sum := 0
	preMax := height[0]
	for i := 1; i < n; i++ {
		if height[i] > preMax {
			preMax = height[i]
		}
		sum += min(preMax, afterMax[i]) - height[i]
	}
	return sum
}

// 方法二：使用常量空间标记最大值，第i个位置能存储的雨水量 = min(leftMax[i], rightMax[i]) - height[i]
func trap2(height []int) int {
	n := len(height)
	left, right, preMax, afterMax := 0, n-1, 0, 0
	sum := 0
	for left < right {
		preMax = max(preMax, height[left])
		afterMax = max(afterMax, height[right])
		if preMax < afterMax {
			sum += preMax - height[left]
			left++
		} else {
			sum += afterMax - height[right]
			right--
		}
	}
	return sum
}

// 7、LeetCode121 买卖股票的最佳时机
/*
给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
示例 1：
输入：[7,1,5,3,6,4]
输出：5
解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
示例 2：
输入：prices = [7,6,4,3,1]
输出：0
解释：在这种情况下, 没有交易完成, 所以最大利润为 0
*/

func TestLeetCode121(t *testing.T) {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
	fmt.Println(maxProfit1([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit1([]int{7, 6, 4, 3, 1}))
	fmt.Println(maxProfit1([]int{2, 1, 4}))
	fmt.Println(maxProfit2([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit2([]int{7, 6, 4, 3, 1}))
	fmt.Println(maxProfit2([]int{2, 1, 4}))
}

// 方法一：暴力法 两次遍历 超时
func maxProfit(prices []int) int {
	res := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			res = max(res, prices[j]-prices[i])
		}
	}
	return res
}

// 方法二：从右往左维护最大值
func maxProfit1(prices []int) int {
	res := 0
	afterMax := 0
	for r := len(prices) - 1; r >= 0; r-- {
		afterMax = max(afterMax, prices[r])
		if prices[r] < afterMax {
			res = max(res, afterMax-prices[r])
		}
	}
	return res
}

// 方法三：从左往右维护最小指
func maxProfit2(prices []int) int {
	res := 0
	preMin := 100000
	for l := 0; l < len(prices); l++ {
		preMin = min(preMin, prices[l])
		if prices[l] > preMin {
			res = max(res, prices[l]-preMin)
		}
	}
	return res
}

// 8、LeetCode209 长度最小的子数组
/*
给定一个含有 n 个正整数的数组和一个正整数 target 。
找出该数组中满足其总和大于等于 target 的长度最小的 子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。
示例 1：
输入：target = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组 [4,3] 是该条件下的长度最小的子数组。
示例 2：
输入：target = 4, nums = [1,4,4]
输出：1
示例 3：
输入：target = 11, nums = [1,1,1,1,1,1,1,1]
输出：0
*/

func TestLeetCode209(t *testing.T) {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	fmt.Println(minSubArrayLen(4, []int{1, 4, 4}))
	fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}))
}

func minSubArrayLen(target int, nums []int) int {
	left := 0
	n := len(nums)
	res := n + 1
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
		for sum >= target {
			res = min(res, i-left+1)
			sum -= nums[left]
			left++
		}
	}
	if res == n+1 {
		res = 0
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type List struct {
	Value int
	Next  *List
}

func TestMergeList(t *testing.T) {
	list1 := &List{
		Value: 1,
		Next: &List{
			Value: 3,
			Next: &List{
				Value: 5,
				Next: &List{
					Value: 8,
				},
			},
		},
	}
	list2 := &List{
		Value: 2,
		Next: &List{
			Value: 4,
			Next: &List{
				Value: 6,
				Next: &List{
					Value: 9,
					Next: &List{
						Value: 10,
						Next: &List{
							Value: 18,
						},
					},
				},
			},
		},
	}
	res := merge(list1, list2)
	fmt.Printf("%+v", res)
}

func merge(l1 *List, l2 *List) *List {
	tmp := &List{0, nil}
	res := tmp
	for l1 != nil && l2 != nil {
		if l1.Value < l2.Value {
			tmp.Next = &List{l1.Value, nil}
			tmp = tmp.Next
			l1 = l1.Next
		} else {
			tmp.Next = &List{l2.Value, nil}
			tmp = tmp.Next
			l2 = l2.Next
		}
	}
	if l1 != nil {
		tmp.Next = l1
	}
	if l2 != nil {
		tmp.Next = l2
	}
	return res.Next
}
