package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	// 解析表单数据
	r.ParseMultipartForm(10 << 20) // 设置最大文件大小为10MB

	// 获取上传的文件
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// 创建目标文件
	dst, err := os.Create(handler.Filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dst.Close()

	// 将上传的文件内容拷贝到目标文件
	if _, err := io.Copy(dst, file); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(w, "文件上传成功: %s", handler.Filename)
}

func downloadFile(w http.ResponseWriter, r *http.Request) {
	// 获取要下载的文件名
	filename := r.URL.Query().Get("file")
	if filename == "" {
		http.Error(w, "文件名不能为空", http.StatusBadRequest)
		return
	}

	// 打开要下载的文件
	file, err := os.Open(filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// 将文件内容作为响应的内容发送给客户端
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/upload", uploadFile)
	http.HandleFunc("/download", downloadFile)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func readFile() {
	// 读取文件内容
	data, err := ioutil.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("File content:")
	fmt.Println(string(data))
}
