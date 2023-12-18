package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

// 实现一个简单的聊天室demo全部代码

// 客户端发送消息
func sendMessage(conn net.Conn) {
	for {
		// 读取用户输入
		msg := readInput()

		// 发送消息到服务器
		_, err := conn.Write([]byte(msg))
		time.Sleep(time.Second)
		if err != nil {
			log.Println("Failed to send message:", err)
			return
		}

		// 如果用户输入"exit"，则退出聊天室
		if strings.TrimSpace(msg) == "exit" {
			fmt.Println("You have left the chat room.")
			os.Exit(0)
		}
	}
}

// 读取用户输入
func readInput() string {
	var msg string
	fmt.Print("Enter message: ")
	fmt.Scanln(&msg)
	return msg
}

// 开启服务器
func startServer() {
	// 监听端口
	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
	defer ln.Close()
	accept, err := ln.Accept()
	if err != nil {
		log.Fatal("Failed to accept connection:", err)
	}
	go receiveMessage(accept)
}

// 服务端接收消息
func receiveMessage(conn net.Conn) {
	for {
		// 读取消息
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			log.Println("Failed to read message:", err)
			return
		}
		// 打印服务器消息
		fmt.Println("Server Received:", string(buffer[:n]))
	}
}

func main() {
	go startServer()
	// 连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal("Failed to connect to server:", err)
	}
	defer conn.Close()

	fmt.Println("Connected to server.")

	// 启动goroutine发送消息
	go sendMessage(conn)

	time.Sleep(time.Minute * 5)
}
