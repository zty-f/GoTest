package algorithm

import (
	"fmt"
	"math/rand"
	"sort"
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

// 方法二：使用第一行和第一列来记录需要置零的行和列，空间复杂度O(1)
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

// 5、LeetCode384 打乱数组
/*
给你一个整数数组 nums ，设计算法来打乱一个没有重复元素的数组。打乱后，数组的所有排列应该是 等可能 的。
实现 Solution class:
Solution(int[] nums) 使用整数数组 nums 初始化对象
int[] reset() 重设数组到它的初始状态并返回
int[] shuffle() 返回数组随机打乱后的结果
示例 1：
输入
["Solution", "shuffle", "reset", "shuffle"]
[[[1, 2, 3]], [], [], []]
输出
[null, [3, 1, 2], [1, 2, 3], [1, 3, 2]]
解释
Solution solution = new Solution([1, 2, 3]);
solution.shuffle();    // 打乱数组 [1,2,3] 并返回结果。任何 [1,2,3]的排列返回的概率应该相同。例如，返回 [3, 1, 2]
solution.reset();      // 重设数组到它的初始状态 [1, 2, 3] 。返回 [1, 2, 3]
solution.shuffle();    // 随机返回数组 [1, 2, 3] 打乱后的结果。例如，返回 [1, 3, 2]
*/

func Test打乱数组(t *testing.T) {
	s := Constructor([]int{1, 2, 3})
	fmt.Println(s.Shuffle())
	fmt.Println(s.Reset())
	fmt.Println(s.Shuffle())
}

type Solution struct {
	origin []int
}

func Constructor(nums []int) Solution {
	return Solution{origin: nums}
}

func (s *Solution) Reset() []int {
	return s.origin
}

// 洗牌算法
/*
洗牌算法的思路很简单。我们有个长度为n的数组nums，对于每个nums[i]来说，都生成一个[i,n−1]范围的随机数，作为random_idx，然后交换nums[i]和nums[random_idx。
为什么说洗牌算法实现的shuffle()返回的数组会有n!种可能呢？
对于nums[0]，它可能会和[0,n−1]范围内的任何一个数交换，有n种可能。
对于nums[1]，它可能会和[1,n−1]范围内的任何一个数交换，有n−1种可能。
...
对于nums[n-1]，它只能和nums[n-1]自己交换，只有1种可能。
所以总的可能性是: n+(n−1)+(n−2)+...+1=n!
*/
func (s *Solution) Shuffle() []int {
	x := make([]int, len(s.origin))
	/*
		copy()函数会将src中的元素逐个复制到dst，不会对切片进行扩容或缩容。
		copy()函数不会创建新的切片，它只是修改目标切片的内容。
	*/
	copy(x, s.origin)
	for i := len(x) - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		x[i], x[j] = x[j], x[i]
	}
	return x
}

// 原生库函数
func (s *Solution) Shuffle1() []int {
	x := make([]int, len(s.origin))
	copy(x, s.origin)
	rand.Shuffle(len(x), func(i, j int) {
		x[i], x[j] = x[j], x[i]
	})
	return x
}

// 6、LeetCode581 最短无序连续子数组
/*
给你一个整数数组 nums ，你需要找出一个 连续子数组 ，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。
请你找出符合题意的 最短 子数组，并输出它的长度。
示例 1：
输入：nums = [2,6,4,8,10,9,15]
输出：5
解释：你只需要对 [6, 4, 8, 10, 9] 进行升序排序，那么整个表都会变为升序排序。
*/
func Test最短无序连续子数组(t *testing.T) {
	fmt.Println(findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15}))
	fmt.Println(findUnsortedSubarray([]int{1, 2, 3, 4}))
	fmt.Println(findUnsortedSubarray1([]int{1, 3, 2, 2, 2}))
	fmt.Println(findUnsortedSubarray1([]int{1, 3, 2, 3, 3}))
}

// 方法一，排序，然后和原数组对比即可
func findUnsortedSubarray(nums []int) int {
	sortNum := numSort(nums)
	l, r := 0, len(nums)-1
	for l < len(nums) && sortNum[l] == nums[l] {
		l++
	}
	for r > l && sortNum[r] == nums[r] {
		r--
	}
	return r - l + 1
}

// 排序
func numSort(nums []int) []int {
	x := make([]int, len(nums))
	copy(x, nums)
	for i := 0; i < len(x); i++ {
		for j := i + 1; j < len(x); j++ {
			if x[j] < x[i] {
				x[i], x[j] = x[j], x[i]
			}
		}
	}
	return x
}

// 方法二：一次遍历，分别维护左右的最大值和最小值
func findUnsortedSubarray1(nums []int) int {
	maxl, minr := -1<<31, 1<<31-1
	l, r := -1, -1
	for i := 0; i < len(nums); i++ {
		// 从左到右维护最大值 右边界之后的数字都将大于左边的所有数字
		if nums[i] < maxl {
			r = i
		} else {
			maxl = nums[i]
		}
		// 从右到左维护最小值 左边界之前的数字都将小于右边的所有数字
		if nums[len(nums)-1-i] > minr {
			l = len(nums) - 1 - i
		} else {
			minr = nums[len(nums)-1-i]
		}
	}
	if r == -1 {
		return 0
	}
	return r - l + 1
}

// 7、LeetCode945 使数组唯一的最小增量
/*
给你一个整数数组 nums 。每次 move 操作将会选择任意一个满足 0 <= i < nums.length 的下标 i，并将 nums[i] 递增 1。
返回使 nums 中的每个值都变成唯一的所需要的最少操作次数。
生成的测试用例保证答案在 32 位整数范围内。
示例 1：
输入：nums = [1,2,2]
输出：1
解释：经过一次 move 操作，数组将变为 [1, 2, 3]。
示例 2：
输入：nums = [3,2,1,2,1,7] 1 2 3 2 3 7
输出：6
解释：经过 6 次 move 操作，数组将变为 [3, 4, 1, 2, 5, 7]。
可以看出 5 次或 5 次以下的 move 操作是不能让数组的每个值唯一的。
*/

func Test使数组唯一的最小增量(t *testing.T) {
	fmt.Println(minIncrementForUnique([]int{1, 2, 2}))
	fmt.Println(minIncrementForUnique1([]int{3, 2, 1, 2, 1, 7}))
	fmt.Println(minIncrementForUnique2([]int{3, 2, 1, 2, 1, 7}))
	fmt.Println(minIncrementForUnique3([]int{3, 2, 1, 2, 1, 7}))
}

// 方法一：排序+贪心
func minIncrementForUnique(nums []int) int {
	sort.Ints(nums)
	res := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			res += nums[i-1] + 1 - nums[i]
			nums[i] = nums[i-1] + 1
		}
	}
	return res
}

// 方法二：计数+贪心
// https://leetcode.cn/problems/minimum-increment-to-make-array-unique/solutions/163214/ji-shu-onxian-xing-tan-ce-fa-onpai-xu-onlogn-yi-ya/
func minIncrementForUnique1(nums []int) int {
	// counter数组统计每个数字的个数。
	// （这里为了防止下面遍历counter的时候每次都走到100000，所以设置了一个max）
	cnt := make([]int, 100002)
	max := 0
	for _, x := range nums {
		cnt[x]++
		if x > max {
			max = x
		}
	}
	res := 0
	// 遍历counter数组，若当前数字的个数cnt大于1个，则只留下1个，其他的cnt-1个后移
	for i := 0; i <= max; i++ {
		if cnt[i] > 1 {
			// 如果某个数出现超过一次，将这些多出的数字都加一，相当于下一个数字加上这些数量即可
			d := cnt[i] - 1
			res += d
			cnt[i+1] += d
		}
	}
	// 最后, counter[max+1]里可能会有从counter[max]后移过来的，counter[max+1]里只留下1个，其它的d个后移。
	// 设 max+1 = x，那么后面的d个数就是[x+1,x+2,x+3,...,x+d],
	// 因此操作次数是[1,2,3,...,d],用求和公式求和。
	d := cnt[max+1] - 1
	res += d * (1 + d) / 2
	return res
}

// 方法三：线性探测法+路径压缩--使用数组
func minIncrementForUnique2(nums []int) int {
	pos := make([]int, 200005)
	for i, _ := range pos {
		pos[i] = -1
	}
	move := 0
	for _, x := range nums {
		move += findPos(pos, x) - x
	}
	return move
}

func findPos(pos []int, x int) int {
	value := pos[x]
	if value == -1 {
		pos[x] = x
		return x
	}
	pos[x] = findPos(pos, value+1)
	return pos[x]
}

// 方法三：线性探测法+路径压缩--使用map
func minIncrementForUnique3(nums []int) int {
	pos := make(map[int]int)
	move := 0
	for _, x := range nums {
		move += findPos2(pos, x) - x
	}
	return move
}

func findPos2(pos map[int]int, x int) int {
	value := pos[x]
	if _, ex := pos[x]; !ex {
		pos[x] = x
		return x
	}
	pos[x] = findPos2(pos, value+1)
	return pos[x]
}
