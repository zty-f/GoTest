package main

import (
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/xuri/excelize/v2"
)

// ExcelExporter Excel导出器
type ExcelExporter struct {
	file      *excelize.File
	outputDir string
}

// NewExcelExporter 创建新的Excel导出器
func NewExcelExporter(outputDir string) *ExcelExporter {
	return &ExcelExporter{
		file:      excelize.NewFile(),
		outputDir: outputDir,
	}
}

// AddBookSheet 添加绘本工作表
func (e *ExcelExporter) AddBookSheet(bookData *APIResponse) error {
	bookTitle := bookData.Data.Ent.BookInfo.Title
	// 清理工作表名称中的特殊字符
	sheetName := cleanSheetName(bookTitle)

	// 创建新的工作表
	index, err := e.file.NewSheet(sheetName)
	if err != nil {
		return fmt.Errorf("创建工作表失败: %v", err)
	}

	// 设置活动工作表
	e.file.SetActiveSheet(index)

	// 设置列宽
	e.file.SetColWidth(sheetName, "A", "A", 8)  // 页码
	e.file.SetColWidth(sheetName, "B", "B", 15) // 绘本内容（图片列）- 2cm约等于9个字符宽度
	e.file.SetColWidth(sheetName, "C", "C", 30) // 英文字幕原文
	e.file.SetColWidth(sheetName, "D", "D", 30) // 中文翻译原文
	e.file.SetColWidth(sheetName, "E", "E", 20) // 英文字母音频

	// 写入绘本信息
	row := 1
	e.file.SetCellValue(sheetName, fmt.Sprintf("A%d", row), "绘本名称")
	e.file.SetCellValue(sheetName, fmt.Sprintf("B%d", row), bookTitle)
	row++

	e.file.SetCellValue(sheetName, fmt.Sprintf("A%d", row), "绘本封面")
	// 尝试插入封面图片
	coverPath := e.getCoverPath(bookTitle)
	if coverPath != "" {
		if err := e.insertImage(sheetName, fmt.Sprintf("B%d", row), coverPath); err != nil {
			// 如果插入图片失败，显示URL
			e.file.SetCellValue(sheetName, fmt.Sprintf("B%d", row), bookData.Data.Ent.BookInfo.Cover.Origin)
			fmt.Printf("插入封面图片失败: %v\n", err)
		}
	} else {
		e.file.SetCellValue(sheetName, fmt.Sprintf("B%d", row), bookData.Data.Ent.BookInfo.Cover.Origin)
	}
	row++

	// 添加空行
	row++

	// 写入表头
	headers := []string{"页码", "绘本内容", "英文字幕原文", "中文翻译原文", "英文字母音频"}
	for i, header := range headers {
		cell := fmt.Sprintf("%c%d", 'A'+i, row)
		e.file.SetCellValue(sheetName, cell, header)
		// 设置表头字体为加粗
		style, err := e.file.NewStyle(&excelize.Style{
			Font: &excelize.Font{
				Bold:   true,
				Italic: true,
				Family: "Times New Roman",
			},
			// 居中
			Alignment: &excelize.Alignment{
				Horizontal: "center",
				Vertical:   "center",
			},
		})
		if err != nil {
			fmt.Println(err)
		}
		e.file.SetCellStyle(sheetName, cell, cell, style)
	}
	row++

	// 写入页面数据
	for _, page := range bookData.Data.Ent.PageInfos {
		// 页码
		e.file.SetCellValue(sheetName, fmt.Sprintf("A%d", row), page.Index)

		// 绘本内容（插入图片）
		imagePath := e.getImagePath(bookTitle, page.Index)
		if imagePath != "" {
			// 插入图片到单元格
			if err := e.insertImage(sheetName, fmt.Sprintf("B%d", row), imagePath); err != nil {
				// 如果插入图片失败，显示提示信息而不是路径
				e.file.SetCellValue(sheetName, fmt.Sprintf("B%d", row), fmt.Sprintf("图片%d (插入失败)", page.Index))
				fmt.Printf("插入图片失败: %v\n", err)
			}
		} else {
			e.file.SetCellValue(sheetName, fmt.Sprintf("B%d", row), fmt.Sprintf("图片%d (未找到)", page.Index))
		}

		// 英文字幕原文
		e.file.SetCellValue(sheetName, fmt.Sprintf("C%d", row), page.ListenText)

		// 中文翻译原文
		e.file.SetCellValue(sheetName, fmt.Sprintf("D%d", row), page.Translation)

		// 英文字母音频（创建超链接）
		audioPath := e.getAudioPath(bookTitle, page.Index)
		if audioPath != "" {
			// 创建超链接 - 使用file://协议指向本地文件
			if err := e.file.SetCellHyperLink(sheetName, fmt.Sprintf("E%d", row), audioPath, "External"); err != nil {
				e.file.SetCellValue(sheetName, fmt.Sprintf("E%d", row), fmt.Sprintf("音频%d (链接失败)", page.Index))
				fmt.Printf("创建音频超链接失败: %v\n", err)
			} else {
				e.file.SetCellValue(sheetName, fmt.Sprintf("E%d", row), fmt.Sprintf("🎵 音频%d", page.Index))
			}
		} else {
			e.file.SetCellValue(sheetName, fmt.Sprintf("E%d", row), fmt.Sprintf("音频%d (未找到)", page.Index))
		}

		row++
	}

	return nil
}

// Save 保存Excel文件
func (e *ExcelExporter) Save(filePath string) error {
	// 删除默认的Sheet1
	e.file.DeleteSheet("Sheet1")

	// 确保目录存在
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("创建目录失败: %v", err)
	}

	// 保存文件
	if err := e.file.SaveAs(filePath); err != nil {
		return fmt.Errorf("保存Excel文件失败: %v", err)
	}

	return nil
}

// cleanSheetName 清理工作表名称中的特殊字符
func cleanSheetName(name string) string {
	// Excel工作表名称限制：最多31个字符，不能包含特殊字符
	replacer := strings.NewReplacer(
		"/", "_",
		"\\", "_",
		":", "_",
		"*", "_",
		"?", "_",
		"[", "_",
		"]", "_",
		" ", "_",
	)

	cleanName := replacer.Replace(name)

	// 限制长度
	if len(cleanName) > 31 {
		cleanName = cleanName[:31]
	}

	return cleanName
}

// getImagePath 获取图片文件路径
func (e *ExcelExporter) getImagePath(bookTitle string, pageIndex int) string {
	cleanTitle := cleanFileName(bookTitle)
	imageDir := filepath.Join(e.outputDir, cleanTitle, "图片")

	// 尝试不同的图片格式
	extensions := []string{".jpg", ".jpeg", ".png", ".gif", ".bmp"}
	for _, ext := range extensions {
		imagePath := filepath.Join(imageDir, fmt.Sprintf("%d%s", pageIndex, ext))
		if _, err := os.Stat(imagePath); err == nil {
			return imagePath
		}
	}
	return ""
}

// getAudioPath 获取音频文件路径
func (e *ExcelExporter) getAudioPath(bookTitle string, pageIndex int) string {
	cleanTitle := cleanFileName(bookTitle)
	audioPath := filepath.Join(e.outputDir, cleanTitle, "音频", fmt.Sprintf("%d.mp3", pageIndex))

	// 检查文件是否存在
	if _, err := os.Stat(audioPath); err == nil {
		return audioPath
	}
	return ""
}

// getCoverPath 获取封面图片路径
func (e *ExcelExporter) getCoverPath(bookTitle string) string {
	cleanTitle := cleanFileName(bookTitle)
	imageDir := filepath.Join(e.outputDir, cleanTitle, "图片")

	// 尝试不同的封面图片格式
	extensions := []string{".jpg", ".jpeg", ".png", ".gif", ".bmp"}
	for _, ext := range extensions {
		coverPath := filepath.Join(imageDir, fmt.Sprintf("封面%s", ext))
		if _, err := os.Stat(coverPath); err == nil {
			return coverPath
		}
	}
	return ""
}

// insertImage 插入图片到Excel单元格
func (e *ExcelExporter) insertImage(sheetName, cell, imagePath string) error {
	// 检查图片文件是否存在
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		return fmt.Errorf("图片文件不存在: %s", imagePath)
	}
	// 设置行高以适应图片
	_, row, err := excelize.CellNameToCoordinates(cell)
	fmt.Printf("成功解析单元格坐标: %s -> 行: %d\n", cell, row)
	if err != nil {
		return fmt.Errorf("解析单元格坐标失败: %v", err)
	}
	e.file.SetRowHeight(sheetName, row, 56) // 设置行高为56.7点，约等于2cm

	fmt.Printf("尝试插入图片: %s\n", imagePath)

	// 使用文档中的标准方法插入图片
	err = e.file.AddPicture(sheetName, cell, imagePath, &excelize.GraphicOptions{
		ScaleX:          1.0,       // 不缩放，让AutoFit控制大小
		ScaleY:          1.0,       // 不缩放，让AutoFit控制大小
		OffsetX:         0,         // 无偏移，完全嵌入
		OffsetY:         0,         // 无偏移，完全嵌入
		LockAspectRatio: false,     // 锁定宽高比
		AutoFit:         true,      // 自动适应单元格大小
		Positioning:     "oneCell", // 固定在单个单元格内
	})

	if err != nil {
		return fmt.Errorf("添加图片失败: %v", err)
	}

	fmt.Printf("成功插入图片: %s\n", imagePath)
	return nil
}
