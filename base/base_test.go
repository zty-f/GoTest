package base

import (
	"fmt"
	"testing"
)

func Test_m1(t *testing.T) {
	m1()
}

type A interface {
	Show()
}

type B struct{}

func (stu *B) Show() {
	fmt.Println("show")
}

func Test1(t *testing.T) {
	var s *B
	if s == nil {
		fmt.Println("s is nil")
	} else {
		fmt.Println("s is not nil")
	}
	var p A = s
	if p == nil {
		fmt.Println("p is nil")
	} else {
		fmt.Println("p is not nil")
	}
	p.Show()
}
