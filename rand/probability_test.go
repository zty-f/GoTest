package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 演示累积概率算法原理
func TestProbability(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	// 设定概率
	highProb := 0.25
	midProb := 0.30
	lowProb := 0.45

	fmt.Println("=== 累积概率算法演示 ===")
	fmt.Printf("高价值概率: %.2f (25%%)\n", highProb)
	fmt.Printf("中价值概率: %.2f (30%%)\n", midProb)
	fmt.Printf("低价值概率: %.2f (45%%)\n\n", lowProb)

	fmt.Println("区间划分：")
	fmt.Printf("[0.00, %.2f) → 高价值\n", highProb)
	fmt.Printf("[%.2f, %.2f) → 中价值\n", highProb, highProb+midProb)
	fmt.Printf("[%.2f, 1.00] → 低价值\n\n", highProb+midProb)

	// 演示几次抽奖
	fmt.Println("=== 单次抽奖演示 ===")
	for i := 0; i < 10; i++ {
		randValue := rand.Float64()
		result := ""

		if randValue < highProb {
			result = "高价值"
		} else if randValue < highProb+midProb {
			result = "中价值"
		} else {
			result = "低价值"
		}

		fmt.Printf("第%2d次: 随机数=%.4f → %s\n", i+1, randValue, result)
	}

	// 大量抽奖验证概率
	fmt.Println("\n=== 大量抽奖验证概率 ===")
	rounds := 100000
	highCount := 0
	midCount := 0
	lowCount := 0

	for i := 0; i < rounds; i++ {
		randValue := rand.Float64()

		if randValue < highProb {
			highCount++
		} else if randValue < highProb+midProb {
			midCount++
		} else {
			lowCount++
		}
	}

	fmt.Printf("总抽奖次数: %d\n", rounds)
	fmt.Printf("高价值: %d次 (实际概率: %.2f%%, 期望: 25%%)\n",
		highCount, float64(highCount)/float64(rounds)*100)
	fmt.Printf("中价值: %d次 (实际概率: %.2f%%, 期望: 30%%)\n",
		midCount, float64(midCount)/float64(rounds)*100)
	fmt.Printf("低价值: %d次 (实际概率: %.2f%%, 期望: 45%%)\n",
		lowCount, float64(lowCount)/float64(rounds)*100)

	// 可视化展示
	fmt.Println("\n=== 概率分布可视化 ===")
	scale := 200 // 每个字符代表200次
	fmt.Printf("高价值 [%.2f]: %s (%d)\n", highProb,
		repeat("█", highCount/scale), highCount)
	fmt.Printf("中价值 [%.2f]: %s (%d)\n", midProb,
		repeat("█", midCount/scale), midCount)
	fmt.Printf("低价值 [%.2f]: %s (%d)\n", lowProb,
		repeat("█", lowCount/scale), lowCount)
}

func repeat(s string, count int) string {
	result := ""
	for i := 0; i < count; i++ {
		result += s
	}
	return result
}
