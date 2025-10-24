package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	// 配置参数
	config := Config{
		BookIDs:   []int64{465442971693328}, // 可以配置多个bookid
		OutputDir: "./output",
		ExcelFile: "./output/绘本数据导出.xlsx",
		APIURL:    "https://sea.pri.ibanyu.com/picturebookopapi/ugc/picturebook/boutique/book/page/get",
		Cookie:    "ipalfish_device_id=e416b0ae2221b84cfa8a84f87b7ace73; _ga=GA1.1.952297019.1743422236; _ga_89WN60ZK2E=GS2.1.s1747893991$o2$g0$t1747893991$j0$l0$h0; did=17483382255380000; utype=op; user=zhangtianyong26331; id=NDAyNTA=; name=5byg5aSp5rOz; phone=ODYtMTM0NTg4MzgyNDg=; logintype=0; lang=zh-cn; groups=WyJkdXdvIiwiNzAwMTg2OTI4IiwicmQiLCI3MDAxNzI3MDIiLCJkZWZhdWx0Il0=; token=275422a3dd89dace07504a935d1d573e",
	}

	// 创建输出目录
	if err := os.MkdirAll(config.OutputDir, 0755); err != nil {
		log.Fatalf("创建输出目录失败: %v", err)
	}

	// 初始化组件
	apiClient := NewAPIClient(config.APIURL, config.Cookie)
	fileDownloader := NewFileDownloader()
	excelExporter := NewExcelExporter(config.OutputDir)

	fmt.Printf("开始处理 %d 个绘本...\n", len(config.BookIDs))

	// 处理每个绘本
	for i, bookID := range config.BookIDs {
		fmt.Printf("\n=== 处理第 %d/%d 个绘本 (ID: %d) ===\n", i+1, len(config.BookIDs), bookID)

		// 获取绘本数据
		fmt.Printf("正在获取绘本数据...\n")
		bookData, err := apiClient.GetBookData(bookID)
		if err != nil {
			log.Printf("获取绘本数据失败 (ID: %d): %v", bookID, err)
			continue
		}

		bookTitle := bookData.Data.Ent.BookInfo.Title
		fmt.Printf("绘本名称: %s\n", bookTitle)

		// 下载文件
		fmt.Printf("正在下载文件...\n")
		downloadResult, err := fileDownloader.DownloadBookFiles(bookData, config.OutputDir)
		if err != nil {
			log.Printf("下载文件失败 (ID: %d): %v", bookID, err)
		} else {
			fmt.Printf("文件下载完成\n")
			fmt.Printf("  - 下载了 %d 张图片\n", len(downloadResult.ImagePaths))
			fmt.Printf("  - 下载了 %d 个音频文件\n", len(downloadResult.AudioPaths))
			if downloadResult.CoverPath != "" {
				fmt.Printf("  - 封面: %s\n", downloadResult.CoverPath)
			}
		}

		// 添加到Excel
		fmt.Printf("正在添加到Excel...\n")
		if err := excelExporter.AddBookSheet(bookData); err != nil {
			log.Printf("添加Excel工作表失败 (ID: %d): %v", bookID, err)
		} else {
			fmt.Printf("Excel工作表添加完成\n")
		}

		// 添加延迟避免请求过于频繁
		if i < len(config.BookIDs)-1 {
			fmt.Printf("等待 2 秒后处理下一个绘本...\n")
			time.Sleep(2 * time.Second)
		}
	}

	// 保存Excel文件
	fmt.Printf("\n正在保存Excel文件...\n")
	if err := excelExporter.Save(config.ExcelFile); err != nil {
		log.Fatalf("保存Excel文件失败: %v", err)
	}

	fmt.Printf("\n=== 处理完成 ===\n")
	fmt.Printf("Excel文件已保存到: %s\n", config.ExcelFile)
	fmt.Printf("文件已下载到: %s\n", config.OutputDir)
}

// loadConfigFromFile 从配置文件加载配置（可选功能）
func loadConfigFromFile(configPath string) (*Config, error) {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("配置文件不存在: %s", configPath)
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %v", err)
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %v", err)
	}

	return &config, nil
}

// saveConfigToFile 保存配置到文件（可选功能）
func saveConfigToFile(config *Config, configPath string) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("序列化配置失败: %v", err)
	}

	if err := os.WriteFile(configPath, data, 0644); err != nil {
		return fmt.Errorf("写入配置文件失败: %v", err)
	}

	return nil
}
