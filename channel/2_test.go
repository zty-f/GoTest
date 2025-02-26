package channel

import (
	"fmt"
	"testing"
	"time"
)

func TestChan1(t *testing.T) {
	// 定义一个通道
	ch := make(chan string)

	// 启动一个 goroutine，向通道中存放数据
	go func() {
		fmt.Println("start goroutine")
		time.Sleep(time.Second * 5)
		ch <- "from 1"
	}()

	msg := <-ch // 从通道中取出数据,会阻塞
	fmt.Println(msg)
}

func TestChanSelect(t *testing.T) {
	// 定义两个通道
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	// 启动两个 goroutine，分别向两个通道中存放数据

	ch1 <- "from 1"

	ch2 <- "from 2"

	// 使用 select 语句非阻塞地从两个通道中获取数据

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	}

	// select会随机的选择一个case来处理，如果有default，且没有case满足条件，就会执行default
}

func counter(out chan<- int) {
	for i := 0; i < 5; i++ {
		out <- i
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}
func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

/*
1、chan<- int是一个只能发送[向通道内加值]的通道，可以发送但是不能接收；

2、<-chan int是一个只能接收[取出通道内的值]的通道，可以接收但是不能发送。

上述方式可以用于比如限制通道在函数中只能发送或只能接收的场景。
*/

func TestChanSingle(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go squarer(ch2, ch1)
	printer(ch2)
}
