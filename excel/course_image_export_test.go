package excel

import (
	"fmt"
	"github.com/spf13/cast"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/xuri/excelize/v2"
)

// CourseLesson 课程课节数据结构
type CourseLesson struct {
	CourseID      string `excel:"course_id"`
	LectureNumber string `excel:"lecture_number"`
	LectureID     string `excel:"lecture_id"`
	CourseName    string `excel:"course_name"`
	LectureName   string `excel:"lecture_name"`
	URI           string `excel:"uri"`
}

// TestExportCourseImages 测试导出课程图片功能
func TestExportCourseImages(t *testing.T) {
	// 读取原始数据
	inputFile := "../export/data.xlsx"
	outputFile := "course_images_export.xlsx"
	outputDir := "./course_images"

	// 读取Excel数据
	var lessons []CourseLesson
	err := NewExcelReader(inputFile).
		SetSheet("查询结果").
		ReadToStruct(&lessons)

	if err != nil {
		t.Fatalf("读取Excel失败: %v", err)
	}

	if len(lessons) == 0 {
		t.Fatal("没有读取到数据")
	}

	// 创建输出目录
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		t.Fatalf("创建输出目录失败: %v", err)
	}

	// 创建新的Excel文件
	f := excelize.NewFile()
	sheetName := "课程数据"
	f.SetSheetName("Sheet1", sheetName)

	// 设置表头（保持所有原始字段）
	headers := []string{"course_id", "lecture_number", "lecture_id", "course_name", "lecture_name", "封面图片"}
	for i, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheetName, cell, header)
	}
	// 设置列宽
	f.SetColWidth(sheetName, "A", "A", 15) // course_id
	f.SetColWidth(sheetName, "B", "B", 15) // lecture_number
	f.SetColWidth(sheetName, "C", "C", 15) // lecture_id
	f.SetColWidth(sheetName, "D", "D", 30) // course_name
	f.SetColWidth(sheetName, "E", "E", 30) // lecture_name
	f.SetColWidth(sheetName, "F", "F", 15) // 封面图片列

	// 创建文件下载器
	downloader := &http.Client{Timeout: 30 * time.Second}

	// 第一步：先按课程分组下载所有图片到文件夹
	// 创建映射：lesson的唯一标识 -> 图片路径
	imagePathMap := make(map[string]string) // key: course_id + lecture_id, value: 图片路径

	// 按课程分组
	courseGroups := make(map[string][]CourseLesson)
	for _, lesson := range lessons {
		courseGroups[lesson.CourseName] = append(courseGroups[lesson.CourseName], lesson)
	}

	t.Logf("开始下载图片，共 %d 个课程", len(courseGroups))

	// 按课程分组下载图片到文件夹
	cnt := 0
	for courseName, courseLessons := range courseGroups {
		if courseName != "Tina家庭阅读营L1" {
			continue
		}
		// 清理课程名称作为文件夹名
		cleanCourseName := cleanFileName(courseName)
		courseDir := filepath.Join(outputDir, cleanCourseName)

		if err := os.MkdirAll(courseDir, 0755); err != nil {
			t.Logf("创建课程文件夹失败: %v", err)
			continue
		}

		// 下载该课程的所有课节图片
		for _, lesson := range courseLessons {
			if lesson.URI == "" {
				continue
			}
			cnt++
			if cnt > 10 {
				break
			}

			// 拼接完整URL
			imageURL := "https://readcamp.cdn.ipalfish.com/" + strings.TrimPrefix(lesson.URI, "/")

			// 清理课节名称作为文件名
			cleanLessonName := cast.ToString(lesson.LectureNumber) + "_" + cleanFileName(lesson.LectureName)
			// 先使用默认扩展名，下载时会根据Content-Type调整
			imagePath := filepath.Join(courseDir, cleanLessonName+".jpg")

			// 下载图片（函数会根据Content-Type自动调整扩展名）
			actualImagePath, err := downloadImageWithPath(downloader, imageURL, imagePath)
			if err != nil {
				t.Logf("下载图片失败 (课程: %s, 课节: %s): %v", courseName, lesson.LectureName, err)
			} else {
				// 创建唯一标识：course_id + lecture_id
				lessonKey := fmt.Sprintf("%s_%s", lesson.CourseID, lesson.LectureID)
				imagePathMap[lessonKey] = actualImagePath
				t.Logf("下载成功: %s -> %s", lesson.LectureName, actualImagePath)
			}
		}
		if cnt > 10 {
			break
		}
	}

	t.Logf("图片下载完成，共下载 %d 张图片", len(imagePathMap))

	// 第二步：处理每一行数据，从已下载的图片中读取并插入Excel
	for rowIndex, lesson := range lessons {
		row := rowIndex + 2 // 从第2行开始（第1行是表头）
		if rowIndex == 10 {
			break
		}

		// 写入所有字段
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), lesson.CourseID)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), lesson.LectureNumber)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), lesson.LectureID)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), lesson.CourseName)
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", row), lesson.LectureName)

		// 从映射中获取图片路径并插入Excel
		if lesson.URI != "" {
			lessonKey := fmt.Sprintf("%s_%s", lesson.CourseID, lesson.LectureID)
			imagePath, exists := imagePathMap[lessonKey]
			if exists {
				// 插入图片到Excel
				if err := insertImageToExcel(f, sheetName, fmt.Sprintf("F%d", row), imagePath); err != nil {
					t.Logf("插入图片失败 (行%d): %v", row, err)
					imageURL := "https://readcamp.cdn.ipalfish.com/" + strings.TrimPrefix(lesson.URI, "/")
					f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), imageURL)
					t.Logf("已设置URL作为备选: %s", imageURL)
				} else {
					// 插入成功，不设置单元格值（图片已插入）
					t.Logf("图片插入成功 (行%d): %s", row, imagePath)
				}
			} else {
				// 如果映射中没有，说明下载失败，显示URL
				imageURL := "https://readcamp.cdn.ipalfish.com/" + strings.TrimPrefix(lesson.URI, "/")
				f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), imageURL)
				t.Logf("图片未下载，设置URL (行%d): %s", row, imageURL)
			}
		} else {
			f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), "")
		}
	}

	// 保存Excel文件
	if err := f.SaveAs(outputFile); err != nil {
		t.Fatalf("保存Excel文件失败: %v", err)
	}

	t.Logf("导出完成！")
	t.Logf("Excel文件: %s", outputFile)
	t.Logf("图片文件夹: %s", outputDir)
}

// downloadImageWithPath 下载图片并返回实际的文件路径
func downloadImageWithPath(client *http.Client, url, filePath string) (string, error) {
	resp, err := client.Get(url)
	if err != nil {
		return "", fmt.Errorf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP错误: %d", resp.StatusCode)
	}

	// 从Content-Type获取图片格式
	contentType := resp.Header.Get("Content-Type")
	ext := getImageExtensionFromContentType(contentType)

	// 如果从Content-Type无法确定，尝试从URL获取
	if ext == "" {
		urlExt := filepath.Ext(url)
		if urlExt != "" {
			ext = urlExt
		} else {
			ext = ".jpg" // 默认使用jpg
		}
	}

	// 如果文件路径没有扩展名或扩展名不匹配，更新文件路径
	actualPath := filePath
	if filepath.Ext(filePath) != ext {
		actualPath = strings.TrimSuffix(filePath, filepath.Ext(filePath)) + ext
	}

	// excelize 不支持 webp 格式，如果检测到 webp，使用 jpg 扩展名（但实际内容仍是 webp，可能会失败）
	// 更好的做法是跳过 webp 或使用图片转换库
	if ext == ".webp" {
		fmt.Printf("警告: 检测到 webp 格式，excelize 可能不支持，将尝试使用 jpg 扩展名: %s\n", actualPath)
		actualPath = strings.TrimSuffix(actualPath, ".webp") + ".jpg"
		ext = ".jpg"
	}

	// 创建文件
	file, err := os.Create(actualPath)
	if err != nil {
		return "", fmt.Errorf("创建文件失败: %v", err)
	}

	// 复制数据
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		file.Close()
		os.Remove(actualPath) // 删除不完整的文件
		return "", fmt.Errorf("写入文件失败: %v", err)
	}

	// 确保数据完全写入磁盘
	if err := file.Sync(); err != nil {
		file.Close()
		os.Remove(actualPath)
		return "", fmt.Errorf("同步文件失败: %v", err)
	}

	// 关闭文件
	if err := file.Close(); err != nil {
		os.Remove(actualPath)
		return "", fmt.Errorf("关闭文件失败: %v", err)
	}

	// 验证下载的文件是否为有效图片
	if !isValidImage(actualPath) {
		os.Remove(actualPath) // 删除无效文件
		return "", fmt.Errorf("下载的文件不是有效的图片格式: %s (Content-Type: %s)", actualPath, contentType)
	}

	// 获取文件信息用于调试
	fileInfo, _ := os.Stat(actualPath)
	fmt.Printf("图片下载成功: %s, 大小: %d bytes, 格式: %s\n", actualPath, fileInfo.Size(), ext)

	return actualPath, nil
}

// getImageExtensionFromContentType 从Content-Type获取图片扩展名
func getImageExtensionFromContentType(contentType string) string {
	contentType = strings.ToLower(contentType)
	switch {
	case strings.Contains(contentType, "image/jpeg") || strings.Contains(contentType, "image/jpg"):
		return ".jpg"
	case strings.Contains(contentType, "image/png"):
		return ".png"
	case strings.Contains(contentType, "image/gif"):
		return ".gif"
	case strings.Contains(contentType, "image/bmp"):
		return ".bmp"
	case strings.Contains(contentType, "image/webp"):
		return ".webp"
	default:
		return ""
	}
}

// isValidImage 检查文件是否为有效的图片
func isValidImage(filePath string) bool {
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

// insertImageToExcel 插入图片到Excel
func insertImageToExcel(f *excelize.File, sheetName, cell, imagePath string) error {
	// 检查图片文件是否存在
	fileInfo, err := os.Stat(imagePath)
	if os.IsNotExist(err) {
		return fmt.Errorf("图片文件不存在: %s", imagePath)
	}
	if err != nil {
		return fmt.Errorf("获取文件信息失败: %v", err)
	}

	// 检查文件大小
	if fileInfo.Size() == 0 {
		return fmt.Errorf("图片文件为空: %s", imagePath)
	}

	// 再次验证文件是否为有效图片
	if !isValidImage(imagePath) {
		return fmt.Errorf("文件不是有效的图片格式: %s", imagePath)
	}

	// 检查文件扩展名是否被 excelize 支持
	ext := strings.ToLower(filepath.Ext(imagePath))
	supportedFormats := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
		".bmp":  true,
	}
	if !supportedFormats[ext] {
		return fmt.Errorf("不支持的图片格式: %s (excelize 仅支持 jpg, png, gif, bmp)", ext)
	}

	// 设置行高以适应图片
	_, row, err := excelize.CellNameToCoordinates(cell)
	fmt.Printf("成功解析单元格坐标: %s -> 行: %d\n", cell, row)
	if err != nil {
		return fmt.Errorf("解析单元格坐标失败: %v", err)
	}
	err = f.SetRowHeight(sheetName, row, 56)
	if err != nil {
		return fmt.Errorf("设置行高失败: %v", err)
	} // 设置行高为56.7点，约等于2cm

	fmt.Printf("尝试插入图片: %s\n", imagePath)

	// 使用文档中的标准方法插入图片（使用相对路径，与export包保持一致）
	err = f.AddPicture(sheetName, cell, imagePath, &excelize.GraphicOptions{
		ScaleX:          1.0,       // 不缩放，让AutoFit控制大小
		ScaleY:          1.0,       // 不缩放，让AutoFit控制大小
		OffsetX:         0,         // 无偏移，完全嵌入
		OffsetY:         0,         // 无偏移，完全嵌入
		LockAspectRatio: false,     // 锁定宽高比（与export包保持一致）
		AutoFit:         true,      // 自动适应单元格大小
		Positioning:     "oneCell", // 固定在单个单元格内
	})

	if err != nil {
		return fmt.Errorf("添加图片失败: %v (路径: %s)", err, imagePath)
	}

	fmt.Printf("成功插入图片: %s\n", imagePath)
	return nil
}

// cleanFileName 清理文件名中的特殊字符
func cleanFileName(filename string) string {
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
