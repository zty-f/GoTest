package channel

import (
	"fmt"
	"testing"
	"time"
)

func Test1(t *testing.T) {
	// 定义两个通道
	ch1 := make(chan string)
	ch2 := make(chan string)

	// 启动两个 goroutine，分别向两个通道中存放数据
	go func() {
		for {
			ch1 <- "from 1"
		}
	}()
	go func() {
		for {
			ch2 <- "from 2"
		}
	}()

	// 使用 select 语句非阻塞地从两个通道中获取数据
	for {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		default:
			// 如果两个通道都没有可用的数据，则执行这里的语句
			fmt.Println("no message received")
		}
	}
}

func Chann(ch chan int, stopCh chan bool) {
	for j := 0; j < 10; j++ {
		ch <- j
		time.Sleep(time.Millisecond * 100)
	}
	stopCh <- true
}

func Test2(t *testing.T) {

	ch := make(chan int)
	stopCh := make(chan bool)

	go Chann(ch, stopCh)

	for {
		select {
		case c := <-ch:
			fmt.Println("Receive C", c)
		case s := <-ch:
			fmt.Println("Receive S", s)
		case _ = <-stopCh:
			goto End
		}
	}
End:
	println("End")
}

// 主协程等待 10 个 goroutine 完成
func Test3(t *testing.T) {

	ch := make(chan int)

	for i := 1; i < 11; i++ {
		i := i
		go func() {
			time.Sleep(time.Second * 2)
			fmt.Println("goroutine", i, "end")
			ch <- i
		}()
	}
	// 等待 10 个 goroutine 完成
	for i := 1; i < 11; i++ {
		fmt.Println("main wait goroutine", <-ch)
	}
	close(ch)
}
