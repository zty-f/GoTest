package csv

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/gocarina/gocsv"
)

// CSVBytesProcessor CSV 二进制流处理器
type CSVBytesProcessor struct {
	// 是否包含标题行
	hasHeader bool
	// 分隔符
	delimiter rune
}

// NewCSVBytesProcessor 创建新的 CSV 二进制流处理器
func NewCSVBytesProcessor() *CSVBytesProcessor {
	return &CSVBytesProcessor{
		hasHeader: true,
		delimiter: ',',
	}
}

// SetHasHeader 设置是否包含标题行
func (c *CSVBytesProcessor) SetHasHeader(hasHeader bool) *CSVBytesProcessor {
	c.hasHeader = hasHeader
	return c
}

// SetDelimiter 设置分隔符
func (c *CSVBytesProcessor) SetDelimiter(delimiter rune) *CSVBytesProcessor {
	c.delimiter = delimiter
	return c
}

// ReadFromBytes 从二进制流读取 CSV 数据到 struct 切片
// data 必须是指向切片的指针，例如: &[]User{}
func (c *CSVBytesProcessor) ReadFromBytes(csvData []byte, data interface{}) error {
	if len(csvData) == 0 {
		return fmt.Errorf("CSV 数据为空")
	}

	// 创建字节读取器
	reader := bytes.NewReader(csvData)

	// 使用 gocsv 从字节读取器读取数据到 struct
	err := gocsv.Unmarshal(reader, data)
	if err != nil {
		return fmt.Errorf("解析 CSV 二进制流失败: %v", err)
	}

	return nil
}

// WriteToBytes 将 struct 切片写入二进制流
// data 必须是切片，例如: []User{}
func (c *CSVBytesProcessor) WriteToBytes(data interface{}) ([]byte, error) {
	// 创建字节缓冲区
	var buf bytes.Buffer

	// 使用 gocsv 将 struct 写入字节缓冲区
	err := gocsv.Marshal(data, &buf)
	if err != nil {
		return nil, fmt.Errorf("写入 CSV 到二进制流失败: %v", err)
	}

	return buf.Bytes(), nil
}

// ReadMapFromBytes 从二进制流读取 CSV 数据到 map 切片
func (c *CSVBytesProcessor) ReadMapFromBytes(csvData []byte) ([]map[string]string, error) {
	if len(csvData) == 0 {
		return nil, fmt.Errorf("CSV 数据为空")
	}

	// 创建字节读取器
	reader := bytes.NewReader(csvData)

	// 使用 gocsv 读取到 map 切片
	var result []map[string]string
	err := gocsv.Unmarshal(reader, &result)
	if err != nil {
		return nil, fmt.Errorf("读取 CSV 二进制流失败: %v", err)
	}

	return result, nil
}

// WriteMapToBytes 将 map 切片写入二进制流
func (c *CSVBytesProcessor) WriteMapToBytes(data []map[string]string) ([]byte, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("数据为空")
	}

	// 创建字节缓冲区
	var buf bytes.Buffer

	// 使用 gocsv 将 map 写入字节缓冲区
	err := gocsv.Marshal(data, &buf)
	if err != nil {
		return nil, fmt.Errorf("写入 CSV 到二进制流失败: %v", err)
	}

	return buf.Bytes(), nil
}

// ReadStringFromBytes 从二进制流读取 CSV 数据到字符串切片
func (c *CSVBytesProcessor) ReadStringFromBytes(csvData []byte) ([][]string, error) {
	if len(csvData) == 0 {
		return nil, fmt.Errorf("CSV 数据为空")
	}

	// 创建字节读取器
	reader := bytes.NewReader(csvData)

	// 使用 gocsv 读取到字符串切片
	var result [][]string
	err := gocsv.Unmarshal(reader, &result)
	if err != nil {
		return nil, fmt.Errorf("读取 CSV 二进制流失败: %v", err)
	}

	return result, nil
}

// WriteStringToBytes 将字符串切片写入二进制流
func (c *CSVBytesProcessor) WriteStringToBytes(data [][]string) ([]byte, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("数据为空")
	}

	// 创建字节缓冲区
	var buf bytes.Buffer

	// 使用 gocsv 将字符串切片写入字节缓冲区
	err := gocsv.Marshal(data, &buf)
	if err != nil {
		return nil, fmt.Errorf("写入 CSV 到二进制流失败: %v", err)
	}

	return buf.Bytes(), nil
}

// ValidateCSVFormat 验证 CSV 二进制流格式是否正确
func (c *CSVBytesProcessor) ValidateCSVFormat(csvData []byte) error {
	if len(csvData) == 0 {
		return fmt.Errorf("CSV 数据为空")
	}

	// 检查是否包含必要的分隔符
	if !bytes.Contains(csvData, []byte(",")) && !bytes.Contains(csvData, []byte(";")) && !bytes.Contains(csvData, []byte("\t")) {
		return fmt.Errorf("CSV 数据格式不正确：缺少分隔符")
	}

	// 检查是否包含换行符（至少应该有一行数据）
	if !bytes.Contains(csvData, []byte("\n")) && !bytes.Contains(csvData, []byte("\r")) {
		return fmt.Errorf("CSV 数据格式不正确：缺少换行符")
	}

	return nil
}

// GetCSVInfo 获取 CSV 二进制流的基本信息
func (c *CSVBytesProcessor) GetCSVInfo(csvData []byte) (map[string]interface{}, error) {
	if len(csvData) == 0 {
		return nil, fmt.Errorf("CSV 数据为空")
	}

	info := make(map[string]interface{})

	// 获取数据大小
	info["size"] = len(csvData)

	// 计算行数
	lines := strings.Split(string(csvData), "\n")
	// 过滤空行
	var nonEmptyLines []string
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			nonEmptyLines = append(nonEmptyLines, line)
		}
	}
	info["lineCount"] = len(nonEmptyLines)

	// 获取第一行作为列数参考
	if len(nonEmptyLines) > 0 {
		columns := strings.Split(nonEmptyLines[0], ",")
		info["columnCount"] = len(columns)
		info["firstLine"] = nonEmptyLines[0]
	}

	// 检查编码（简单检查是否包含中文字符）
	if bytes.Contains(csvData, []byte("中")) || bytes.Contains(csvData, []byte("文")) {
		info["encoding"] = "可能包含中文"
	} else {
		info["encoding"] = "ASCII"
	}

	return info, nil
}

// ConvertToCSVString 将二进制流转换为 CSV 字符串
func (c *CSVBytesProcessor) ConvertToCSVString(csvData []byte) (string, error) {
	if len(csvData) == 0 {
		return "", fmt.Errorf("CSV 数据为空")
	}

	return string(csvData), nil
}

// ConvertFromCSVString 将 CSV 字符串转换为二进制流
func (c *CSVBytesProcessor) ConvertFromCSVString(csvString string) []byte {
	return []byte(csvString)
}

// GetDelimiter 获取当前设置的分隔符
func (c *CSVBytesProcessor) GetDelimiter() rune {
	return c.delimiter
}

// GetHasHeader 获取是否包含标题行
func (c *CSVBytesProcessor) GetHasHeader() bool {
	return c.hasHeader
}
