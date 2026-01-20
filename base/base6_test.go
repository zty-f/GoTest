package base

import (
	"fmt"
	"testing"
	"time"
)

func TestNowYear(t *testing.T) {
	year := time.Now().Year() - 1
	fmt.Printf("%d\n", year)
	fmt.Println(int32(year))
}
