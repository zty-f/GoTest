package excel

import (
	"fmt"
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

	// 按课程分组
	courseGroups := make(map[string][]CourseLesson)
	for _, lesson := range lessons {
		courseGroups[lesson.CourseName] = append(courseGroups[lesson.CourseName], lesson)
	}

	// 处理每一行数据
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

		// 下载并插入图片到uri列（F列）
		if lesson.URI != "" {
			// 拼接完整URL
			imageURL := "https://readcamp.cdn.ipalfish.com/" + strings.TrimPrefix(lesson.URI, "/")

			// 下载图片到临时文件
			ext := filepath.Ext(imageURL)
			if ext == "" {
				ext = ".jpg"
			}
			tempImagePath := filepath.Join(outputDir, fmt.Sprintf("temp_%d%s", rowIndex, ext))
			if err := downloadImage(downloader, imageURL, tempImagePath); err != nil {
				t.Logf("下载图片失败 (行%d): %v, URL: %s", row, err, imageURL)
				f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), imageURL)
			} else {
				// 插入图片到Excel
				if err := insertImageToExcel(f, sheetName, fmt.Sprintf("F%d", row), tempImagePath); err != nil {
					t.Logf("插入图片失败 (行%d): %v", row, err)
					f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), imageURL)
				}
				// 删除临时文件
				// os.Remove(tempImagePath)
			}
		} else {
			f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), "")
		}
	}

	// 按课程分组下载图片到文件夹
	// for courseName, courseLessons := range courseGroups {
	// 	// 清理课程名称作为文件夹名
	// 	cleanCourseName := cleanFileName(courseName)
	// 	courseDir := filepath.Join(outputDir, cleanCourseName)
	//
	// 	if err := os.MkdirAll(courseDir, 0755); err != nil {
	// 		t.Logf("创建课程文件夹失败: %v", err)
	// 		continue
	// 	}

	// 下载该课程的所有课节图片
	// for _, lesson := range courseLessons {
	// 	if lesson.URI == "" {
	// 		continue
	// 	}
	//
	// 	// 拼接完整URL
	// 	imageURL := "https://readcamp.cdn.ipalfish.com/" + strings.TrimPrefix(lesson.URI, "/")
	//
	// 	// 清理课节名称作为文件名
	// 	cleanLessonName := cleanFileName(lesson.LectureName)
	// 	imagePath := filepath.Join(courseDir, cleanLessonName+".jpg")
	//
	// 	// 下载图片
	// 	if err := downloadImage(downloader, imageURL, imagePath); err != nil {
	// 		t.Logf("下载图片失败 (课程: %s, 课节: %s): %v", courseName, lesson.LectureName, err)
	// 	} else {
	// 		t.Logf("下载成功: %s -> %s", lesson.LectureName, imagePath)
	// 	}
	// }
	// }

	// 保存Excel文件
	if err := f.SaveAs(outputFile); err != nil {
		t.Fatalf("保存Excel文件失败: %v", err)
	}

	t.Logf("导出完成！")
	t.Logf("Excel文件: %s", outputFile)
	t.Logf("图片文件夹: %s", outputDir)
}

// downloadImage 下载图片
func downloadImage(client *http.Client, url, filePath string) error {
	resp, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP错误: %d", resp.StatusCode)
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

	return nil
}

// insertImageToExcel 插入图片到Excel
func insertImageToExcel(f *excelize.File, sheetName, cell, imagePath string) error {
	// 设置行高以适应图片
	_, row, err := excelize.CellNameToCoordinates(cell)
	fmt.Printf("成功解析单元格坐标: %s -> 行: %d\n", cell, row)
	if err != nil {
		return fmt.Errorf("解析单元格坐标失败: %v", err)
	}
	f.SetRowHeight(sheetName, row, 56) // 设置行高为56.7点，约等于2cm

	fmt.Printf("尝试插入图片: %s\n", imagePath)

	// 使用文档中的标准方法插入图片
	err = f.AddPicture(sheetName, cell, imagePath, &excelize.GraphicOptions{
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
