package usage

import (
	"codeup.aliyun.com/61e54b0e0bb300d827e1ae27/backend/golib/logger"
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"runtime/debug"
	"sync"
	"syscall"
	"testing"
	"time"
)

/*
channel 的一些应用场景包括不限于下面的这些：

1.同步，用来在多个协程之间的同步操作，如同一时间，只能有一个协程工作，其他协程只能等待，这里也可以对应到同步互斥的场景
2.数据传递，可通过 channel 在多个协程间进行数据传递，比如在 生产者-消费者 模式中使用，生产者向 channel 发送数据，消费者从 channel 中消费数据，也就涉及到对 channel 写-读 操作
3.管道，channel 作为管道，可将数据从一处转移到其他处，在数据流中常见
4.缓冲，在有缓存的 channel 中发送数据，没有读取的情况下，就可以将数据缓存起来，有助于缓解读者压力
5.信号通知，通过 channel 的读写完成信号通知
6.范围迭代，通过 for val := range chan 的操作，当遇到 channel 关闭的时候就会退出迭代，从而完成 channel 数据所有接收
7.并发控制，通过 channel 的有缓存类型，限制一定的并发数量，不至于无限制并发消耗系统资源
8.定时器，有很多标准库的结构体都带有 channel 功能，比如 time.Timer，就是一个定时器，通过定时器定期执行某些任务
*/

// 通过限制 limit 的缓存数量，决定并发时有多少协程在并行运行
var limit = make(chan struct{}, 3)

// 1 并发控制
func Test1(t *testing.T) {
	tasks := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	for i, v := range tasks {
		// 每个task开启一个协程
		go func(i, v int) {
			// 通过chan控制并发
			limit <- struct{}{}

			// 具体的任务执行
			fmt.Printf("Time: %v, Goroutine: %v exec i : %d, v: %v\n", time.Now().Format(time.RFC3339Nano), i, i, v)
			time.Sleep(time.Second)

			<-limit
		}(i, v)
	}

	time.Sleep(6 * time.Second)
	fmt.Printf("Time: %v, 主线程退出！", time.Now().Format(time.RFC3339))
}

// 2 管道 | 范围迭代 | 数据传输
func Test2(t *testing.T) {
	numCh := make(chan int)

	waitGroup := &sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		waitGroup.Add(1)
		go sender(numCh, waitGroup)
	}

	// 等待子协程执行完成才能关闭管道
	go SafeGo(context.Background(), func() {
		waitGroup.Wait()
		close(numCh)
	})

	// 这里会阻塞，numCh 是无缓冲通道，子协程写入数据后需要有人读取后才会继续执行
	for val := range numCh { // for range 读取channel里面的数据的时候，如果channel没有数据了，就会阻塞，直到close通道才会停止
		fmt.Println("Main recv val: ", val)
	}

	fmt.Println("Data transform model Done.")

}

// 仅仅是对函数本身做Panic Recover，自身不会启动协程，需要在协程中调用
func SafeGo(ctx context.Context, fun func()) {
	defer func(ctx context.Context) {
		if err := recover(); err != nil {
			stack := string(debug.Stack())
			logger.Ex(ctx, "SafeGo", "Goroutine Recover: %+v, stack is %s", err, stack)
		}
	}(ctx)
	fun()
}

func sender(num chan int, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	fmt.Println("Send goroutine start...")
	randomNum := rand.Intn(100)
	num <- randomNum
	fmt.Println("Send val to chan:", randomNum)
	time.Sleep(time.Second)
}

// 3 数据传递 -> 生产者-消费者模型
func Test3(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)

	numCh := make(chan int, 1)
	go producer(numCh, &wg)
	go consumer(numCh, &wg)

	wg.Wait()
	fmt.Println("本次生产者-消费者模型结束.")
}

func producer(num chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	count := 0
	for count < 1 {
		for i := 0; i < 3; i++ {
			randomNum := rand.Intn(100)
			num <- randomNum
			fmt.Printf("生产者持续发送消息：%d\n", randomNum)
		}
		time.Sleep(time.Second)
		count++
	}

	close(num)
	fmt.Println("生产者关闭通道channel.")
}

func consumer(num chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range num {
		fmt.Printf("消费者持续接收消息：%d\n", val)
		time.Sleep(time.Second)
	}

	fmt.Println("消费者接收完成.")
}

// 4 互斥同步
func Test4(t *testing.T) {
	mu := NewMutexLock()
	ok := mu.tryLock()
	fmt.Printf("locked v %v\n", ok) // true
	ok = mu.tryLock()
	fmt.Printf("locked v %v\n", ok) // false，需要等待持锁方释放后才能再获取锁
}

type mutex struct {
	ch chan struct{}
}

func NewMutexLock() *mutex {
	// 定义 1 个 cap 的互斥锁，拿到数据读取到数据就是拿到锁，发送成功数据就是解锁
	mu := &mutex{ch: make(chan struct{}, 1)}
	mu.ch <- struct{}{}
	return mu
}

// 读取到数据就是锁定状态
func (m *mutex) lock() {
	<-m.ch
}

// 写入数据就是解锁状态
func (m *mutex) unlock() {
	select {
	case m.ch <- struct{}{}:
	default:
		panic("unlock of unlocked mutex")
	}
}

func (m *mutex) tryLock() bool {
	select {
	case <-m.ch:
		return true
	default:
	}

	return false
}

// 锁超时
func (m *mutex) lockTimeout(timeout time.Duration) bool {
	timer := time.NewTimer(timeout)
	select {
	case <-m.ch:
		timer.Stop()
		return true
	case <-timer.C:
	}

	return false
}

func (m *mutex) isLocked() bool {
	return len(m.ch) == 0
}

// 5 信号通知
func Test5(t *testing.T) {
	var (
		closing = make(chan struct{})
		closed  = make(chan struct{})
	)

	go func() {
		// 模拟业务处理
		for {
			select {
			case <-closing:
				return
			default: // 只要没有打断程序运行，这里就会一直执行下去，直到收到打断信号
				fmt.Println("处理业务中...")
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	// 处理 ctrl + c 等中断信号
	terminalChan := make(chan os.Signal)
	signal.Notify(terminalChan, syscall.SIGINT, syscall.SIGTERM) // 待接收信号写入 terminalChan
	<-terminalChan                                               // 收到信号，读取信号

	close(closing) // 关闭 closing，发送信号给业务处理 goroutine，让其结束业务

	// 因为已经结束业务，这里做一些清理工作
	go doCleanup(closed) // 清理结束后，会关闭 closed chan

	// 主线程监听 closed 的关闭情况
	select {
	case <-closed:
	case <-time.After(time.Second):
		fmt.Println("超时了，不等cleanup了。")
	}

	fmt.Println("优雅退出。")
}

// 清理资源占用
func doCleanup(closed chan struct{}) {
	time.Sleep(5 * time.Second) // 模拟实际的清理
	close(closed)               // 关闭 chan
}

// 6 定时器
func Test6(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println("timer start...")
	<-timer.C
	fmt.Println("timer end...")
}
