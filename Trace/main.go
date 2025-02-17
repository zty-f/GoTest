package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

/*
运行程序

$ go run trace.go
Hello World
会得到一个 trace.out 文件，然后我们可以用一个工具打开，来分析这个文件。

$ go tool trace trace.out
2025/02/17 18:55:04 Parsing trace...
2025/02/17 18:55:04 Splitting trace...
2025/02/17 18:55:04 Opening browser. Trace viewer is listening on http://127.0.0.1:52026
*/

func main() {

	// 创建trace文件
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	// 启动trace goroutine
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	// main
	fmt.Println("Hello World")
}
