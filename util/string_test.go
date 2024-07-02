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
