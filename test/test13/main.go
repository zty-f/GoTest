package main

//func main() {
//	arr := []int{1, 2, 3, 4, 5}
//	for i := range arr {
//		fmt.Println(i)
//	}
//
//	for i, v := range arr {
//		fmt.Println(i)
//		fmt.Println(v)
//	}
//
//	fmt.Println(util.Now().Unix())
//
//	for i := 0; i < 5; i++ {
//		for j := 0; j < 5; j++ {
//			fmt.Printf("%d %d\n", i, j)
//			if j == 2 {
//				break
//			}
//		}
//	}
//
//}

//func main() {
//	var wg sync.WaitGroup
//	foo := make(chan int)
//	bar := make(chan int)
//	wg.Add(1)
//	go func() {
//		defer wg.Done()
//		select {
//		case foo <- <-bar:
//		default:
//			println("default")
//		}
//	}()
//	wg.Wait()
//}
import (
	"fmt"
	"net/http"
	"time"
)

//func main() {
//	ch := make(chan int)
//	go func() {
//		select {
//		case ch <- getVal(1):
//			fmt.Println("in first case")
//		case ch <- getVal(2):
//			fmt.Println("in second case")
//		default:
//			fmt.Println("default")
//		}
//	}()
//
//	fmt.Println("The val:", <-ch)
//}

func main() {
	ch := make(chan int, 10)

	go func() {
		var i = 1
		for {
			time.Sleep(time.Second * 1)
			i++
			ch <- i
		}
	}()

	go func() {
		// http 监听8080, 开启 pprof
		if err := http.ListenAndServe(":8080", nil); err != nil {
			fmt.Println("listen failed")
		}
	}()

	for {
		select {
		case x := <-ch:
			println(x)
		case <-time.After(2 * time.Second):
			println(time.Now().Unix())
		}
	}
}

//func main() {
//	ch1 := make(chan string, 1)
//
//	go func() {
//		util.Sleep(util.Second * 2)
//		ch1 <- "hello"
//	}()
//
//	select {
//	case res := <-ch1:
//		fmt.Println(res)
//	case t := <-util.After(util.Second * 1):
//		fmt.Println("timeout", t)
//	}
//}

func getVal(i int) int {
	fmt.Println("getVal, i=", i)
	return i
}
