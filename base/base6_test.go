package base

import (
	"fmt"
	"testing"
	"time"
)

type CBD struct {
	StartTs int64 `json:"start_ts"`
}

func TestNowYear(t *testing.T) {
	year := time.Now().Year() - 1
	fmt.Printf("%d\n", year)
	fmt.Println(int32(year))
	_ = "{\"start_ts\":1750003200,\"guide_exp_limit\":250,\"practice_exp_limit\":300,\"practice_exp_per_min\":1,\"guide_exp_once\":25,\"re_guide_exp_limit\":50,\"re_guide_exp_once\":5,\"lian_yi_lian_limit\":100,\"re_lian_yi_lian_limit\":50,\"lian_yi_lian_once\":10,\"re_lian_yi_lian_once\":5,\"follow_read_limit\":100,\"re_follow_read_limit\":50,\"follow_read_once\":10,\"re_follow_read_once\":5,\"max_rank_cert_template_uri\":\"10/sea/ec/dc/53fbb18524000db301ff04863151\"}"

	var cbd CBD
	x := &cbd
	fmt.Println(cbd)
	fmt.Println(x.StartTs)
}

// 开关状态
type SwitchStatus int

const (
	Off  SwitchStatus = iota // 关
	On                       // 开
	None                     // 无
)

func (s SwitchStatus) String() string {
	switch s {
	case On:
		return "开"
	case Off:
		return "关"
	case None:
		return "无"
	default:
		return "未知"
	}
}

// GetFinalStatus 根据全局开关和个人开关计算最终结果
// 逻辑规则：
// 1. 个人开关优先：如果个人开关是"开"或"关"，则直接使用个人开关的值
// 2. 个人开关为"无"时，使用全局开关的值，但全局开关为"无"时，结果为"关"
func GetFinalStatus(globalSwitch, personalSwitch SwitchStatus) SwitchStatus {
	// 个人开关不为"无"时，直接返回个人开关状态
	if personalSwitch != None {
		return personalSwitch
	}

	// 个人开关为"无"时，根据全局开关决定
	// 全局开关为"开"时返回"开"，否则返回"关"
	if globalSwitch == On {
		return On
	}
	return Off
}

func TestSwitchLogic(t *testing.T) {
	testCases := []struct {
		globalSwitch   SwitchStatus
		personalSwitch SwitchStatus
		expected       SwitchStatus
		description    string
	}{
		{On, On, On, "全局开关：开，个人开关：开 -> 开"},
		{On, Off, Off, "全局开关：开，个人开关：关 -> 关"},
		{On, None, On, "全局开关：开，个人开关：无 -> 开"},
		{Off, On, On, "全局开关：关，个人开关：开 -> 开"},
		{Off, Off, Off, "全局开关：关，个人开关：关 -> 关"},
		{Off, None, Off, "全局开关：关，个人开关：无 -> 关"},
		{None, On, On, "全局开关：无，个人开关：开 -> 开"},
		{None, Off, Off, "全局开关：无，个人开关：关 -> 关"},
		{None, None, Off, "全局开关：无，个人开关：无 -> 关"},
	}

	fmt.Println("\n开关逻辑测试:")
	fmt.Println("| 全局开关 | 个人开关 | 期望结果 | 实际结果 | 状态 |")
	fmt.Println("|---------|---------|---------|---------|------|")

	allPassed := true
	for _, tc := range testCases {
		actual := GetFinalStatus(tc.globalSwitch, tc.personalSwitch)
		status := "✓"
		if actual != tc.expected {
			status = "✗"
			allPassed = false
			t.Errorf("%s: 期望 %s, 实际 %s", tc.description, tc.expected, actual)
		}
		fmt.Printf("| %s     | %s     | %s     | %s     | %s    |\n",
			tc.globalSwitch, tc.personalSwitch, tc.expected, actual, status)
	}

	if allPassed {
		fmt.Println("\n✓ 所有测试通过!")
	} else {
		fmt.Println("\n✗ 部分测试失败")
	}
}
