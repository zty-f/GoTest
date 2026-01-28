package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// GiftValue 奖品价值等级
type GiftValue int

const (
	HighValue GiftValue = 1 // 高价值
	MidValue  GiftValue = 2 // 中价值
	LowValue  GiftValue = 3 // 低价值
)

// RewardInfo 奖品信息
type RewardInfo struct {
	GiftID    int       `json:"gift_id"`
	GiftValue GiftValue `json:"gift_value"`
}

// LotteryRuleInfo 抽奖规则信息
type LotteryRuleInfo struct {
	Tag           string  `json:"tag"`
	HighValueProb float64 `json:"high_value_prob"` // 高价值概率
	MidValueProb  float64 `json:"mid_value_prob"`  // 中价值概率
	LowValueProb  float64 `json:"low_value_prob"`  // 低价值概率
}

// UserInfo 用户信息
type UserInfo struct {
	UserID string   `json:"user_id"`
	Tags   []string `json:"tags"` // 用户的标签列表
}

// LotteryConfig 抽奖配置
type LotteryConfig struct {
	RewardInfos      []RewardInfo      `json:"reward_infos"`
	LotteryRuleInfos []LotteryRuleInfo `json:"lottery_rule_infos"`
}

// LotterySystem 抽奖系统
type LotterySystem struct {
	config *LotteryConfig
	random *rand.Rand
}

// NewLotterySystem 创建抽奖系统
func NewLotterySystem(config *LotteryConfig) *LotterySystem {
	return &LotterySystem{
		config: config,
		random: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// DrawLottery 通用抽奖方法
// 参数：
//   - user: 用户信息
//
// 
// 返回：
//   - *RewardInfo: 抽中的奖品，如果没有抽中返回nil
//   - error: 错误信息
func (ls *LotterySystem) DrawLottery(user *UserInfo) (*RewardInfo, error) {
	// 1. 根据用户标签匹配抽奖规则（取第一个匹配的规则）
	rule := ls.findMatchingRule(user)
	if rule == nil {
		return nil, fmt.Errorf("no matching lottery rule found for user: %s", user.UserID)
	}

	// 2. 根据规则的概率确定奖品价值等级
	giftValue := ls.determineGiftValue(rule)

	// 3. 从对应价值等级的奖品池中随机抽取一个奖品
	reward := ls.selectRewardByValue(giftValue)
	if reward == nil {
		return nil, fmt.Errorf("no reward found for gift value: %d", giftValue)
	}

	return reward, nil
}

// findMatchingRule 查找匹配的抽奖规则
// 遍历规则列表，返回第一个与用户标签匹配的规则
func (ls *LotterySystem) findMatchingRule(user *UserInfo) *LotteryRuleInfo {
	for _, rule := range ls.config.LotteryRuleInfos {
		for _, userTag := range user.Tags {
			if rule.Tag == userTag {
				return &rule
			}
		}
	}
	return nil
}

// determineGiftValue 根据概率确定奖品价值等级
// 使用加权随机算法
func (ls *LotterySystem) determineGiftValue(rule *LotteryRuleInfo) GiftValue {
	// 生成0-1之间的随机数
	randValue := ls.random.Float64()

	// 累积概率判断
	if randValue < rule.HighValueProb {
		return HighValue
	} else if randValue < rule.HighValueProb+rule.MidValueProb {
		return MidValue
	} else {
		return LowValue
	}
}

// selectRewardByValue 从指定价值等级的奖品中随机选择一个
func (ls *LotterySystem) selectRewardByValue(giftValue GiftValue) *RewardInfo {
	// 筛选出对应价值等级的所有奖品
	var candidates []RewardInfo
	for _, reward := range ls.config.RewardInfos {
		if reward.GiftValue == giftValue {
			candidates = append(candidates, reward)
		}
	}

	// 如果没有对应等级的奖品，返回nil
	if len(candidates) == 0 {
		return nil
	}

	// 随机选择一个奖品
	idx := ls.random.Intn(len(candidates))
	return &candidates[idx]
}

// DrawLotteryWithDetail 带详细信息的抽奖方法
// 返回抽奖过程的详细信息，便于调试和日志记录
func (ls *LotterySystem) DrawLotteryWithDetail(user *UserInfo) (reward *RewardInfo, rule *LotteryRuleInfo, giftValue GiftValue, err error) {
	// 1. 匹配规则
	rule = ls.findMatchingRule(user)
	if rule == nil {
		err = fmt.Errorf("no matching lottery rule found for user: %s", user.UserID)
		return
	}

	// 2. 确定价值等级
	giftValue = ls.determineGiftValue(rule)

	// 3. 选择奖品
	reward = ls.selectRewardByValue(giftValue)
	if reward == nil {
		err = fmt.Errorf("no reward found for gift value: %d", giftValue)
		return
	}

	return
}

// ---------- 测试用例 ----------

func TestLotterySystem(t *testing.T) {
	// 创建测试配置
	config := &LotteryConfig{
		RewardInfos: []RewardInfo{
			{GiftID: 30023, GiftValue: HighValue},
			{GiftID: 30024, GiftValue: MidValue},
			{GiftID: 30025, GiftValue: LowValue},
		},
		LotteryRuleInfos: []LotteryRuleInfo{
			{
				Tag:           "high",
				HighValueProb: 0.25,
				MidValueProb:  0.3,
				LowValueProb:  0.45,
			},
			{
				Tag:           "middle",
				HighValueProb: 0.3,
				MidValueProb:  0.45,
				LowValueProb:  0.25,
			},
			{
				Tag:           "low",
				HighValueProb: 0.2,
				MidValueProb:  0.4,
				LowValueProb:  0.4,
			},
		},
	}

	// 创建抽奖系统
	lotterySystem := NewLotterySystem(config)

	// 测试不同标签的用户
	testCases := []struct {
		name string
		user *UserInfo
	}{
		{
			name: "High Tag User",
			user: &UserInfo{UserID: "user001", Tags: []string{"high"}},
		},
		{
			name: "Middle Tag User",
			user: &UserInfo{UserID: "user002", Tags: []string{"middle"}},
		},
		{
			name: "Low Tag User",
			user: &UserInfo{UserID: "user003", Tags: []string{"low"}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			reward, err := lotterySystem.DrawLottery(tc.user)
			if err != nil {
				t.Errorf("DrawLottery failed: %v", err)
				return
			}
			t.Logf("User %s drew gift: ID=%d, Value=%d", tc.user.UserID, reward.GiftID, reward.GiftValue)
		})
	}
}

func TestLotterySystemWithDetail(t *testing.T) {
	config := &LotteryConfig{
		RewardInfos: []RewardInfo{
			{GiftID: 30023, GiftValue: HighValue},
			{GiftID: 30024, GiftValue: MidValue},
			{GiftID: 30025, GiftValue: LowValue},
		},
		LotteryRuleInfos: []LotteryRuleInfo{
			{
				Tag:           "high",
				HighValueProb: 0.25,
				MidValueProb:  0.3,
				LowValueProb:  0.45,
			},
		},
	}

	lotterySystem := NewLotterySystem(config)
	user := &UserInfo{UserID: "user001", Tags: []string{"high"}}

	reward, rule, giftValue, err := lotterySystem.DrawLotteryWithDetail(user)
	if err != nil {
		t.Fatalf("DrawLotteryWithDetail failed: %v", err)
	}

	t.Logf("User: %s", user.UserID)
	t.Logf("Matched Rule: %s (High:%.2f, Mid:%.2f, Low:%.2f)",
		rule.Tag, rule.HighValueProb, rule.MidValueProb, rule.LowValueProb)
	t.Logf("Gift Value: %d", giftValue)
	t.Logf("Reward: ID=%d, Value=%d", reward.GiftID, reward.GiftValue)
}

// TestLotteryProbability 测试概率分布
func TestLotteryProbability(t *testing.T) {
	config := &LotteryConfig{
		RewardInfos: []RewardInfo{
			{GiftID: 30023, GiftValue: HighValue},
			{GiftID: 30024, GiftValue: MidValue},
			{GiftID: 30025, GiftValue: LowValue},
		},
		LotteryRuleInfos: []LotteryRuleInfo{
			{
				Tag:           "high",
				HighValueProb: 0.25,
				MidValueProb:  0.3,
				LowValueProb:  0.45,
			},
		},
	}

	lotterySystem := NewLotterySystem(config)
	user := &UserInfo{UserID: "user001", Tags: []string{"high"}}

	// 进行大量抽奖以验证概率分布
	rounds := 10000
	highCount := 0
	midCount := 0
	lowCount := 0

	for i := 0; i < rounds; i++ {
		reward, err := lotterySystem.DrawLottery(user)
		if err != nil {
			t.Fatalf("DrawLottery failed: %v", err)
		}

		switch reward.GiftValue {
		case HighValue:
			highCount++
		case MidValue:
			midCount++
		case LowValue:
			lowCount++
		}
	}

	// 计算实际概率
	highProb := float64(highCount) / float64(rounds)
	midProb := float64(midCount) / float64(rounds)
	lowProb := float64(lowCount) / float64(rounds)

	t.Logf("Total rounds: %d", rounds)
	t.Logf("High Value: count=%d, prob=%.4f (expected=0.25)", highCount, highProb)
	t.Logf("Mid Value: count=%d, prob=%.4f (expected=0.30)", midCount, midProb)
	t.Logf("Low Value: count=%d, prob=%.4f (expected=0.45)", lowCount, lowProb)

	// 验证概率是否在合理范围内（允许5%的误差）
	tolerance := 0.05
	if abs(highProb-0.25) > tolerance {
		t.Errorf("High value probability out of range: got %.4f, expected 0.25±%.2f", highProb, tolerance)
	}
	if abs(midProb-0.30) > tolerance {
		t.Errorf("Mid value probability out of range: got %.4f, expected 0.30±%.2f", midProb, tolerance)
	}
	if abs(lowProb-0.45) > tolerance {
		t.Errorf("Low value probability out of range: got %.4f, expected 0.45±%.2f", lowProb, tolerance)
	}
}

// 辅助函数：计算绝对值
func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

// TestMultipleRewardsPerValue 测试每个价值等级有多个奖品的情况
func TestMultipleRewardsPerValue(t *testing.T) {
	config := &LotteryConfig{
		RewardInfos: []RewardInfo{
			{GiftID: 30023, GiftValue: HighValue},
			{GiftID: 30026, GiftValue: HighValue},
			{GiftID: 30024, GiftValue: MidValue},
			{GiftID: 30027, GiftValue: MidValue},
			{GiftID: 30025, GiftValue: LowValue},
			{GiftID: 30028, GiftValue: LowValue},
		},
		LotteryRuleInfos: []LotteryRuleInfo{
			{
				Tag:           "test",
				HighValueProb: 1.0, // 100%抽中高价值，用于测试多奖品随机性
				MidValueProb:  0.0,
				LowValueProb:  0.0,
			},
		},
	}

	lotterySystem := NewLotterySystem(config)
	user := &UserInfo{UserID: "user001", Tags: []string{"test"}}

	// 统计每个奖品被抽中的次数
	giftCount := make(map[int]int)
	rounds := 1000

	for i := 0; i < rounds; i++ {
		reward, err := lotterySystem.DrawLottery(user)
		if err != nil {
			t.Fatalf("DrawLottery failed: %v", err)
		}
		giftCount[reward.GiftID]++
	}

	t.Logf("Gift distribution after %d rounds:", rounds)
	for giftID, count := range giftCount {
		t.Logf("Gift %d: %d times (%.2f%%)", giftID, count, float64(count)/float64(rounds)*100)
	}

	// 验证两个高价值奖品都有被抽中
	if giftCount[30023] == 0 || giftCount[30026] == 0 {
		t.Errorf("Not all high value gifts were drawn")
	}
}

// Example 示例：如何使用抽奖系统
func ExampleLotterySystem() {
	// 1. 准备配置
	config := &LotteryConfig{
		RewardInfos: []RewardInfo{
			{GiftID: 30023, GiftValue: HighValue},
			{GiftID: 30024, GiftValue: MidValue},
			{GiftID: 30025, GiftValue: LowValue},
		},
		LotteryRuleInfos: []LotteryRuleInfo{
			{
				Tag:           "vip",
				HighValueProb: 0.5,
				MidValueProb:  0.3,
				LowValueProb:  0.2,
			},
		},
	}

	// 2. 创建抽奖系统
	lotterySystem := NewLotterySystem(config)

	// 3. 用户抽奖
	user := &UserInfo{
		UserID: "user123",
		Tags:   []string{"vip"},
	}

	reward, err := lotterySystem.DrawLottery(user)
	if err != nil {
		fmt.Printf("抽奖失败: %v\n", err)
		return
	}

	fmt.Printf("恭喜！您抽中了奖品 %d (价值等级: %d)\n", reward.GiftID, reward.GiftValue)
}
