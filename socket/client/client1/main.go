package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	// 客户端可以发送单行数据到服务器，然后就退出
	reader := bufio.NewReader(os.Stdin) // os.Stdin是标准输入
	for {
		// 从终端读取一行用户输入并准备发送给服务器
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
		}
		// 如果用户输入的是exit，则退出程序
		line = strings.Trim(line, " \r\n")
		if line == "exit" {
			fmt.Println("Exiting...")
			conn.Close()
			break
		}
		// 再将line发送给服务器
		_, err = conn.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Println("conn.Write err=", err)
		}

		// 读取服务器返回的数据
		msg := make([]byte, 128)
		n, err := conn.Read(msg)
		if err != nil {
			fmt.Println("Error reading from server:", err)
			return
		}
		fmt.Println("Received from server:", string(msg[:n]))
	}
}
