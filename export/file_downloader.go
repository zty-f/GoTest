package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// FileDownloader 文件下载器
type FileDownloader struct {
	Client *http.Client
}

// NewFileDownloader 创建新的文件下载器
func NewFileDownloader() *FileDownloader {
	return &FileDownloader{
		Client: &http.Client{
			Timeout: 60 * time.Second,
		},
	}
}

// DownloadFile 下载文件
func (d *FileDownloader) DownloadFile(url, filePath string) error {
	// 创建目录
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("创建目录失败: %v", err)
	}

	// 检查文件是否已存在
	if _, err := os.Stat(filePath); err == nil {
		fmt.Printf("文件已存在，跳过下载: %s\n", filePath)
		return nil
	}

	// 创建请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("创建请求失败: %v", err)
	}

	// 设置User-Agent
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36")

	// 发送请求
	resp, err := d.Client.Do(req)
	if err != nil {
		return fmt.Errorf("下载文件失败: %v", err)
	}
	defer resp.Body.Close()

	// 检查响应状态
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("下载失败，状态码: %d", resp.StatusCode)
	}

	// 创建文件
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("创建文件失败: %v", err)
	}
	defer file.Close()

	// 复制数据
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("写入文件失败: %v", err)
	}

	fmt.Printf("下载完成: %s\n", filePath)
	return nil
}

// DownloadBookFiles 下载绘本的所有文件
func (d *FileDownloader) DownloadBookFiles(bookData *APIResponse, outputDir string) error {
	bookTitle := bookData.Data.Ent.BookInfo.Title
	// 清理文件名中的特殊字符
	cleanTitle := cleanFileName(bookTitle)

	// 创建绘本目录
	bookDir := filepath.Join(outputDir, cleanTitle)
	imageDir := filepath.Join(bookDir, "图片")
	audioDir := filepath.Join(bookDir, "音频")

	// 创建目录
	if err := os.MkdirAll(imageDir, 0755); err != nil {
		return fmt.Errorf("创建图片目录失败: %v", err)
	}
	if err := os.MkdirAll(audioDir, 0755); err != nil {
		return fmt.Errorf("创建音频目录失败: %v", err)
	}

	// 下载封面
	coverURL := bookData.Data.Ent.BookInfo.Cover.Origin
	if coverURL != "" {
		coverPath := filepath.Join(imageDir, "封面.jpg")
		if err := d.DownloadFile(coverURL, coverPath); err != nil {
			fmt.Printf("下载封面失败: %v\n", err)
		}
	}

	// 下载每页的文件
	for _, page := range bookData.Data.Ent.PageInfos {
		// 下载背景图片
		if page.BGPicture != "" {
			imagePath := filepath.Join(imageDir, fmt.Sprintf("%d.jpg", page.Index))
			if err := d.DownloadFile(page.BGPicture, imagePath); err != nil {
				fmt.Printf("下载第%d页图片失败: %v\n", page.Index, err)
			}
		}

		// 下载音频文件
		if page.ListenAudioURL != "" {
			audioPath := filepath.Join(audioDir, fmt.Sprintf("%d.mp3", page.Index))
			if err := d.DownloadFile(page.ListenAudioURL, audioPath); err != nil {
				fmt.Printf("下载第%d页音频失败: %v\n", page.Index, err)
			}
		}
	}

	return nil
}

// cleanFileName 清理文件名中的特殊字符
func cleanFileName(filename string) string {
	// 替换Windows和Unix不允许的字符
	replacer := strings.NewReplacer(
		"/", "_",
		"\\", "_",
		":", "_",
		"*", "_",
		"?", "_",
		"\"", "_",
		"<", "_",
		">", "_",
		"|", "_",
		" ", "_",
	)
	return replacer.Replace(filename)
}
