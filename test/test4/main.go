package main

import (
	"fmt"
	"sort"
)

func main() {
	a1 := []string{"a", "4", "c"}
	a2 := []string{"d", "4", "f"}
	fmt.Println(intersection(a1, a2))
}

func intersection(arr1 []string, arr2 []string) bool {
	sort.Strings(arr1)
	sort.Strings(arr2)
	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] == arr2[j] {
			return true
		} else if arr1[i] < arr2[j] {
			i++
		} else {
			j++
		}
	}
	return false
}
