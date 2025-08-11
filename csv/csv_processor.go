package csv

import (
	"byte
	"fmt"
	"os"


	"github.com/gocarina/gocsv"
)

// CSVProcessor CSV 文件处理器
type CSVProcessor struct {
	// 文件路径
	filePath string
	// 是否包含标题行
	hasHeader bool
	// 分隔符
	delimiter rune
}

// NewCSVProcessor 创建新的 CSV 处理器
func NewCSVProcessor(filePath string) *CSVProcessor {
	return &CSVProcessor{
		filePath:  filePath,
		hasHeader: true,
		delimiter: ',',
	}
}

 ',',
	}
}

// NewCSVProcessorFromBytes 创建新的 CSV 处理器（用于处理二进制流）
func NewCSVProcessorFromBytes() *CSVProcessor {
	return &CSVProcessor{
		hasHeader: true,
		delimiter: ',',
	}
}

// SetHasHeader 设置是否包含标题行
func (c *CSVProces
// SetHasHeader 设置是否包含标题行
func (c *CSVProcessor) SetHasHeader(hasHeader bool) *CSVProcessor {
	c.hasHeader = hasHeader
	return c
}

SetDelimiter(delimiter rune) *CSVProcessor {
	c.delimiter = delimiter
	return c
}

// ReadFromBytes 从二进制流读取 CSV 数据到 struct 切片
// data 必须是指向切片的指针，例如: &[]User{}
func (c *CSVProcessor) ReadFromBytes(csvData []byte, data interface{}) error {
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
func (c *CSVProcessor) WriteToBytes(data interface{}) ([]byte, error) {
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
func (c *CSVProcessor) ReadMapFromBytes(csvData []byte) ([]map[string]string, error) {
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
func (c *CSVProcessor) WriteMapToBytes(data []map[string]string) ([]byte, error) {
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
func (c *CSVProcessor) ReadStringFromBytes(csvData []byte) ([][]string, error) {
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
func (c *CSVProcessor) WriteStringToBytes(data [][]string) ([]byte, error) {
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
func (c *CSVProcessor) ValidateCSVFormat(csvData []byte) error {
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
func (c *CSVProcessor) GetCSVInfo(csvData []byte) (map[string]interface{}, error) {
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
		
// SetDelimiter 设置分隔符
func (c *CSVProcessor) SetDelimiter(delimiter rune) *CSVProcessor {
	c.delimiter = delimiter
	return c
}

// ReadToStruct 读取 CSV 文件到 struct 切片
// data 必须是指向切片的指针，例如: &[]User{}
func (c *CSVProcessor) ReadToStruct(data interface{}) error {
	file, err := os.OpenFile(c.filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return fmt.Errorf("打开文件失败: %v", err)
	}
	defer file.Close()

	// 使用 gocsv 读取数据到 struct
	err = gocsv.UnmarshalFile(file, data)
	if err != nil {
		return fmt.Errorf("解析 CSV 失败: %v", err)
	}

	return nil
}

// WriteFromStruct 将 struct 切片写入 CSV 文件
// data 必须是切片，例如: []User{}
func (c *CSVProcessor) WriteFromStruct(data interface{}) error {
	file, err := os.OpenFile(c.filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return fmt.Errorf("创建文件失败: %v", err)
	}
	defer file.Close()

	// 使用 gocsv 将 struct 写入 CSV
	err = gocsv.MarshalFile(data, file)
	if err != nil {
		return fmt.Errorf("写入 CSV 失败: %v", err)
	}

	return nil
}

// AppendToStruct 追加数据到现有 CSV 文件
func (c *CSVProcessor) AppendToStruct(data interface{}) error {
	// 检查文件是否存在
	var existingData interface{}
	if c.FileExists() {
		// 读取现有数据
		file, err := os.OpenFile(c.filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
		if err != nil {
			return fmt.Errorf("打开现有文件失败: %v", err)
		}
		defer file.Close()

		// 这里需要根据具体类型来读取，暂时简化处理
		// 实际使用时可能需要更复杂的逻辑
		_ = existingData
	}

	// 打开文件进行追加
	file, err := os.OpenFile(c.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return fmt.Errorf("打开文件进行追加失败: %v", err)
	}
	defer file.Close()

	// 使用 gocsv 追加数据
	err = gocsv.MarshalFile(data, file)
	if err != nil {
		return fmt.Errorf("追加数据失败: %v", err)
	}

	return nil
}

// ReadToMap 读取 CSV 文件到 map 切片
func (c *CSVProcessor) ReadToMap() ([]map[string]string, error) {
	file, err := os.OpenFile(c.filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("打开文件失败: %v", err)
	}
	defer file.Close()

	// 使用 gocsv 读取到 map 切片
	var result []map[string]string
	err = gocsv.UnmarshalFile(file, &result)
	if err != nil {
		return nil, fmt.Errorf("读取 CSV 失败: %v", err)
	}

	return result, nil
}

// WriteFromMap 将 map 切片写入 CSV 文件
func (c *CSVProcessor) WriteFromMap(data []map[string]string) error {
	if len(data) == 0 {
		return fmt.Errorf("数据为空")
	}

	file, err := os.OpenFile(c.filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return fmt.Errorf("创建文件失败: %v", err)
	}
	defer file.Close()

	// 使用 gocsv 将 map 写入 CSV
	err = gocsv.MarshalFile(data, file)
	if err != nil {
		return fmt.Errorf("写入 CSV 失败: %v", err)
	}

	return nil
}

// ReadToString 读取 CSV 文件到字符串切片
func (c *CSVProcessor) ReadToString() ([][]string, error) {
	file, err := os.OpenFile(c.filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("打开文件失败: %v", err)
	}
	defer file.Close()

	// 使用 gocsv 读取到字符串切片
	var result [][]string
	err = gocsv.UnmarshalFile(file, &result)
	if err != nil {
		return nil, fmt.Errorf("读取 CSV 失败: %v", err)
	}

	return result, nil
}

// WriteFromString 将字符串切片写入 CSV 文件
func (c *CSVProcessor) WriteFromString(data [][]string) error {
	if len(data) == 0 {
		return fmt.Errorf("数据为空")
	}

	file, err := os.OpenFile(c.filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return fmt.Errorf("创建文件失败: %v", err)
	}
	defer file.Close()

	// 使用 gocsv 将字符串切片写入 CSV
	err = gocsv.MarshalFile(data, file)
	if err != nil {
		return fmt.Errorf("写入 CSV 失败: %v", err)
	}

	return nil
}

// GetFilePath 获取文件路径
func (c *CSVProcessor) GetFilePath() string {
	return c.filePath
}

// FileExists 检查文件是否存在
func (c *CSVProcessor) FileExists() bool {
	_, err := os.Stat(c.filePath)
	return err == nil
}

// DeleteFile 删除文件
func (c *CSVProcessor) DeleteFile() error {
	if c.FileExists() {
		return os.Remove(c.filePath)
	}
	return nil
}

// GetFileSize 获取文件大小
func (c *CSVProcessor) GetFileSize() (int64, error) {
	info, err := os.Stat(c.filePath)
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}

// 设置 gocsv 的全局配置
func init() {
	// 设置 CSV 读取器
	gocsv.SetCSVReader(gocsv.LazyCSVReader)
}
