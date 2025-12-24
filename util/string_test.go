package util

import (
	"fmt"
	"testing"
)

func TestCutStringAfterSep(t *testing.T) {
	str := "测试【限时奖励】你好"
	sep := "【"
	got := CutStringAfterSep(str, sep)
	fmt.Println(got)
}

func TestGenerateRandomFloatBetween(t *testing.T) {
	min := 1
	max := 3
	got := GenerateRandomFloatBetween(min, max)
	fmt.Println(got)
}
