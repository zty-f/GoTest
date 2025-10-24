package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/xuri/excelize/v2"
)

// ExcelExporter Excel导出器
type ExcelExporter struct {
	file *excelize.File
}

// NewExcelExporter 创建新的Excel导出器
func NewExcelExporter() *ExcelExporter {
	return &ExcelExporter{
		file: excelize.NewFile(),
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

	// 写入绘本信息
	row := 1
	e.file.SetCellValue(sheetName, fmt.Sprintf("A%d", row), "绘本名称")
	e.file.SetCellValue(sheetName, fmt.Sprintf("B%d", row), bookTitle)
	row++

	e.file.SetCellValue(sheetName, fmt.Sprintf("A%d", row), "绘本封面")
	e.file.SetCellValue(sheetName, fmt.Sprintf("B%d", row), bookData.Data.Ent.BookInfo.Cover.Origin)
	row++

	// 添加空行
	row++

	// 写入表头
	headers := []string{"页码", "绘本内容", "英文字幕原文", "中文翻译原文", "英文字母音频"}
	for i, header := range headers {
		cell := fmt.Sprintf("%c%d", 'A'+i, row)
		e.file.SetCellValue(sheetName, cell, header)
	}
	row++

	// 写入页面数据
	for _, page := range bookData.Data.Ent.PageInfos {
		// 页码
		e.file.SetCellValue(sheetName, fmt.Sprintf("A%d", row), page.Index)

		// 绘本内容（背景图片URL）
		e.file.SetCellValue(sheetName, fmt.Sprintf("B%d", row), page.BGPicture)

		// 英文字幕原文
		e.file.SetCellValue(sheetName, fmt.Sprintf("C%d", row), page.ListenText)

		// 中文翻译原文
		e.file.SetCellValue(sheetName, fmt.Sprintf("D%d", row), page.Translation)

		// 英文字母音频
		e.file.SetCellValue(sheetName, fmt.Sprintf("E%d", row), page.ListenAudioURL)

		row++
	}

	// 设置列宽
	e.file.SetColWidth(sheetName, "A", "A", 10) // 页码
	e.file.SetColWidth(sheetName, "B", "B", 50) // 绘本内容
	e.file.SetColWidth(sheetName, "C", "C", 30) // 英文字幕原文
	e.file.SetColWidth(sheetName, "D", "D", 30) // 中文翻译原文
	e.file.SetColWidth(sheetName, "E", "E", 50) // 英文字母音频

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
