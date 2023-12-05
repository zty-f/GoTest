package main

import (
	"testing"
	"time"
)

func BenchmarkTraverseMap(b *testing.B) {
	time.Sleep(time.Second * 1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TraverseMap()
	}
}

func BenchmarkTraverseArray(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TraverseArray()
	}
}

func TestTraverseMap(t *testing.T) {
	TraverseMap()
}
