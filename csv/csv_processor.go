package csv

import (
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

// SetHasHeader 设置是否包含标题行
func (c *CSVProcessor) SetHasHeader(hasHeader bool) *CSVProcessor {
	c.hasHeader = hasHeader
	return c
}

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
