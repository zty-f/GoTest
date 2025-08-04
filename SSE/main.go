package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func sseHandler(w http.ResponseWriter, r *http.Request) {
	// 设置 SSE 所需的响应头
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Cache-Control")

	// 创建一个通道用于发送数据
	messageChan := make(chan string)

	// 模拟数据生成
	go func() {
		for i := 1; i <= 5; i++ {
			message := fmt.Sprintf("消息 #%d - 时间: %s", i, time.Now().Format("15:04:05"))
			messageChan <- message
			time.Sleep(2 * time.Second)
		}
		close(messageChan)
	}()

	// 监听客户端断开连接
	go func() {
		<-r.Context().Done()
		log.Println("客户端断开连接")
	}()

	// 发送数据流
	for message := range messageChan {
		// SSE 数据格式：data: 消息内容\n\n
		fmt.Fprintf(w, "data: %s\n\n", message)

		// 刷新缓冲区，确保数据立即发送
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}
	}
}

func main() {
	http.HandleFunc("/events", sseHandler)

	// 提供静态文件服务（HTML页面）
	http.Handle("/", http.FileServer(http.Dir(".")))

	log.Println("服务器启动在 http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
