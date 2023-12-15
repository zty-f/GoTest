package util

import (
	"fmt"
	"testing"
	"time"
)

func TestSub(t *testing.T) {
	now := time.Now()
	todayStartTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	tn := time.Now()
	fmt.Println(tn.Sub(todayStartTime))
	fmt.Println(tn.Sub(todayStartTime) < 10*time.Minute)
	fmt.Println(todayStartTime.Sub(tn))
	fmt.Println(todayStartTime.Sub(tn) < 10*time.Minute)
}
