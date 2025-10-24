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
func (d *FileDownloader) DownloadBookFiles(bookData *APIResponse, outputDir string) (*DownloadResult, error) {
	bookTitle := bookData.Data.Ent.BookInfo.Title
	// 清理文件名中的特殊字符
	cleanTitle := cleanFileName(bookTitle)

	// 创建绘本目录
	bookDir := filepath.Join(outputDir, cleanTitle)
	imageDir := filepath.Join(bookDir, "图片")
	audioDir := filepath.Join(bookDir, "音频")

	// 创建目录
	if err := os.MkdirAll(imageDir, 0755); err != nil {
		return nil, fmt.Errorf("创建图片目录失败: %v", err)
	}
	if err := os.MkdirAll(audioDir, 0755); err != nil {
		return nil, fmt.Errorf("创建音频目录失败: %v", err)
	}

	// 初始化下载结果
	result := &DownloadResult{
		BookTitle:  bookTitle,
		ImagePaths: make(map[int]string),
		AudioPaths: make(map[int]string),
	}

	// 下载封面
	coverURL := bookData.Data.Ent.BookInfo.Cover.Origin
	if coverURL != "" {
		coverPath := filepath.Join(imageDir, "封面.jpg")
		if err := d.DownloadFile(coverURL, coverPath); err != nil {
			fmt.Printf("下载封面失败: %v\n", err)
		} else {
			result.CoverPath = coverPath
		}
	}

	// 下载每页的文件
	for _, page := range bookData.Data.Ent.PageInfos {
		// 下载背景图片
		if page.BGPicture != "" {
			imagePath := filepath.Join(imageDir, fmt.Sprintf("%d.jpg", page.Index))
			if err := d.DownloadFile(page.BGPicture, imagePath); err != nil {
				fmt.Printf("下载第%d页图片失败: %v\n", page.Index, err)
			} else {
				// 检查下载的文件是否为有效图片
				if d.isValidImage(imagePath) {
					result.ImagePaths[page.Index] = imagePath
				} else {
					fmt.Printf("第%d页图片格式无效，尝试其他扩展名\n", page.Index)
					// 尝试其他扩展名
					extensions := []string{".png", ".jpeg", ".gif", ".bmp"}
					for _, ext := range extensions {
						newPath := filepath.Join(imageDir, fmt.Sprintf("%d%s", page.Index, ext))
						if err := d.DownloadFile(page.BGPicture, newPath); err == nil && d.isValidImage(newPath) {
							result.ImagePaths[page.Index] = newPath
							// 删除原来的jpg文件
							os.Remove(imagePath)
							break
						}
					}
				}
			}
		}

		// 下载音频文件
		if page.ListenAudioURL != "" {
			// 检测音频文件格式
			audioExt := d.getAudioExtension(page.ListenAudioURL)
			audioPath := filepath.Join(audioDir, fmt.Sprintf("%d%s", page.Index, audioExt))
			if err := d.DownloadFile(page.ListenAudioURL, audioPath); err != nil {
				fmt.Printf("下载第%d页音频失败: %v\n", page.Index, err)
			} else {
				result.AudioPaths[page.Index] = audioPath
			}
		}
	}

	return result, nil
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

// getAudioExtension 从URL中检测音频文件扩展名
func (d *FileDownloader) getAudioExtension(url string) string {
	// 从URL中提取文件扩展名
	ext := filepath.Ext(url)

	// 如果没有扩展名或扩展名不是音频格式，默认为.mp3
	if ext == "" {
		return ".mp3"
	}

	// 检查是否为支持的音频格式
	supportedAudioFormats := map[string]bool{
		".mp3":  true,
		".wav":  true,
		".m4a":  true,
		".aac":  true,
		".ogg":  true,
		".flac": true,
		".wma":  true,
	}

	// 转换为小写进行比较
	extLower := strings.ToLower(ext)
	if supportedAudioFormats[extLower] {
		return extLower
	}

	// 如果不支持的格式，默认为.mp3
	return ".mp3"
}

// isValidImage 检查文件是否为有效的图片
func (d *FileDownloader) isValidImage(filePath string) bool {
	file, err := os.Open(filePath)
	if err != nil {
		return false
	}
	defer file.Close()

	// 读取文件头部的几个字节来检测图片格式
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return false
	}

	// 检查常见的图片格式的文件头
	contentType := http.DetectContentType(buffer)
	return strings.HasPrefix(contentType, "image/")
}

// DownloadResult 下载结果
type DownloadResult struct {
	BookTitle  string
	ImagePaths map[int]string // 页码 -> 图片路径
	AudioPaths map[int]string // 页码 -> 音频路径
	CoverPath  string         // 封面路径
}
