package algorithm

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

func TestFourSum(t *testing.T) {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 2))
	fmt.Println(fourSum([]int{2, 2, 2, 2}, 8))
}

// 1. [18]四数之和
func fourSum(nums []int, target int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)
	n := len(nums)
	for a := 0; a < n-3; a++ {
		x := nums[a]
		if a > 0 && x == nums[a-1] { // 跳过重复数字
			continue
		}
		if x+nums[a+1]+nums[a+2]+nums[a+3] > target {
			break
		}
		if x+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}
		for b := a + 1; b < n-2; b++ {
			y := nums[b]
			if b > a+1 && y == nums[b-1] { // 跳过重复数字
				continue
			}
			if x+y+nums[b+1]+nums[b+2] > target {
				break
			}
			if x+y+nums[n-2]+nums[n-1] < target {
				continue
			}
			c := b + 1
			d := n - 1
			for c < d {
				if x+y+nums[c]+nums[d] < target {
					c++
				} else if x+y+nums[c]+nums[d] > target {
					d--
				} else {
					res := []int{x, y, nums[c], nums[d]}
					result = append(result, res)
					for c++; c < d && nums[c] == nums[c-1]; c++ {
					} // 跳过重复数字
					for d--; d > c && nums[d] == nums[d+1]; d-- {
					} // 跳过重复数字
				}
			}
		}
	}
	return result
}

/*
问题描述
在一个游戏中，小W拥有 n 个英雄，每个英雄的初始能力值均为 1。她可以通过升级操作来提升英雄的能力值，最多可以进行 k 次升级。
https://www.marscode.cn/practice/9ejn6or69o3rdn?problem_id=7414004855076470828
游戏规则：
每个英雄的奖励只能获得一次
升级操作的选择是自由的，可以多次选择同一个英雄进行升级
请计算在最多进行 k 次升级操作后，小W能获得的最大奖励总和。
*/

func solution(n, k int, b, c []int) int {
	// 获取b最大能力值
	maxB := 0
	for i := 0; i < n; i++ {
		if b[i] > maxB {
			maxB = b[i]
		}
	}
	// 初始化数组d,d[i]表示能力值到i的最小升级次数
	d := make([]int, maxB+1)
	for i := 1; i <= maxB; i++ {
		if i == 1 {
			d[i] = 0 // 初始时，d[1] = 0，表示能力值1不需要升级。
		} else {
			d[i] = int(math.Pow(9, 10)) // 初始时，i > 1时初始化为一个较大的数字，后续通过动态规划更新。 go语言里面 9^10 是按位运算 = 3
		}
	}
	for i := 1; i < maxB+1; i++ {
		for j := 1; j <= i; j++ {
			t := i + i/j   // 以能力值i为基础进行一次升级
			if t <= maxB { // 如果升级后的能力值小于等于maxB代表升级后在我们所需的范围里面
				d[t] = int(math.Min(float64(d[t]), float64(d[i]+1)))
			}
		}
	}
	// 更新数组b，b[i]表示当前所处的能力值所需的最小升级次数
	upGradeSum := 0
	for i := 0; i < n; i++ {
		b[i] = d[b[i]]
		upGradeSum += b[i]
	}
	// 如果所有能力值达到所需的升级次数小于等于可以升级的次数k，则可以获取全部的奖励
	if upGradeSum <= k {
		res := 0
		for _, v := range c {
			res += v
		}
		return res
	}

	// 如果能够升级的次数有限，需要通过背包算法获取最优的奖励获取方式
	// 定义数组dp[i]表示升级次数为i时的最大奖励
	f := make([]int, k+1)

	// 遍历每个英雄的升级成本和奖励
	for i := 0; i < n; i++ {
		cost := b[i]
		weight := c[i]

		// 从 k 到 cost-1 进行倒序遍历
		for j := k; j >= cost; j-- {
			// 更新 f[j] 的值，取当前值和 f[j-cost] + weight 的最大值
			f[j] = max(f[j], f[j-cost]+weight)
		}
	}

	// 返回 f[k] 的值，即在 k 次升级操作内能获得的最大奖励总和
	return f[k]
}
func Test4(t *testing.T) {
	// 测试用例 1
	fmt.Println(solution(4, 4, []int{1, 7, 5, 2}, []int{2, 6, 5, 2}) == 9)

	// 测试用例 2
	fmt.Println(solution(3, 0, []int{3, 5, 2}, []int{5, 4, 7}) == 0)

	// 测试用例 3
	fmt.Println(solution(3, 3, []int{3, 5, 2}, []int{5, 4, 7}) == 12)

	// 测试用例 4  0 1 2 2 3
	fmt.Println(solution(5, 5, []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}) == 10)

	// 测试用例 5
	fmt.Println(solution(5, 10, []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}) == 15)

	// 测试用例 6
	fmt.Println(solution(5, 10, []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}) == 15)

	// 测试用例 7
	fmt.Println(solution(5, 10, []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}) == 15)

	// 测试用例 8
	fmt.Println(solution(5, 10, []int{5, 4, 3, 2, 1}, []int{5, 4, 3, 2, 1}) == 15)

	// 测试用例 9
	fmt.Println(solution(5, 10, []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}) == 15)

	// 测试用例 10
	fmt.Println(solution(5, 10, []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}) == 15)
}
