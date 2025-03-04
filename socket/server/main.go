package main

// 网络编程server端
import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	fmt.Println("服务器开始监听...")
	// net.Listen("tcp", "0.0.0.0:8888")
	// 1.tcp 表示使用网络协议是 tcp
	// 2. 0.0.0.0 表示监听 8888 端口
	listener, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("服务器监听失败:", err)
		return
	}
	defer listener.Close()

	for {
		// 等待客户端连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("服务器接收客户端连接失败:", err)
			continue
		} else {
			fmt.Println("服务器接收客户端连接成功:", conn.RemoteAddr())
		}
		// 处理客户端请求 多个协程处理-多个连接处理互不影响
		go process(conn)
	}
}

func process(conn net.Conn) {
	// 处理客户端请求
	defer conn.Close() // 关闭连接，如果不关闭会因为连接数不释放，导致服务器无法接收新的连接
	for {
		// 创建新的切片
		buf := make([]byte, 1024)
		// 读取客户端数据。如果客户端不 conn.Write()，conn.Read()会一直阻塞
		fmt.Printf("等待客户端发送数据...%v\n", conn.RemoteAddr())
		n, err := conn.Read(buf)

		if err == io.EOF {
			fmt.Println("客户端连接断开")
			break
		}
		if string(buf[:n]) == "exit\n" {
			fmt.Println("客户端请求退出")
			os.Exit(0)
		}
		// 打印客户端发送的数据
		fmt.Printf("客户端%s发送的数据:%v", conn.RemoteAddr().String(), string(buf[:n])) // [:n]一定要写，否则会打印多余的东西出来
		fmt.Println("模拟数据处理中·········")
		time.Sleep(time.Second * 5)
		fmt.Println("模拟数据处理完成·········")
		// 回复客户端
		_, err = conn.Write([]byte(fmt.Sprintf("已收到%s数据\n", conn.RemoteAddr().String())))
		if err != nil {
			fmt.Println("回复客户端失败:", err)
			break
		}
	}
}
