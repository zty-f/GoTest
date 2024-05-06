package algorithm

import (
	"fmt"
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
