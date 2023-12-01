package main

import "fmt"

func f(n int) (r int) {
	defer func() {
		r += n
		recover()
	}()
	var f func()
	defer f()
	f = func() {
		r += 2
	}
	return n + 1
}
func main() {
	fmt.Println(f(3))
}

//func main() {
//	var m = map[string]int{
//		"A": 21,
//		"B": 22,
//		"C": 23,
//	}
//	counter := 0
//	for k, v := range m {
//		if counter == 0 {
//			delete(m, "A")
//		}
//		counter++
//		fmt.Println(k, v)
//	}
//	fmt.Println("counter is ", counter)
//}
