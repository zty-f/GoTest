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
