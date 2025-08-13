package csv

import (
	"bytes"
	"fmt"

	"github.com/gocarina/gocsv"
)

// CSVBytesProcessor2 CSV 二进制流处理器
type CSVBytesProcessor2 struct {
	// 是否包含标题行
	hasHeader bool
	// 分隔符
	delimiter rune
	// 跳过的行数
	skipRowsBegin int
	skipRowsEnd   int
}

// NewCSVBytesProcessor2 创建新的 CSV 二进制流处理器
func NewCSVBytesProcessor2() *CSVBytesProcessor2 {
	return &CSVBytesProcessor2{
		hasHeader:     true,
		delimiter:     ',',
		skipRowsBegin: 0,
		skipRowsEnd:   0,
	}
}

// SetHasHeader 设置是否包含标题行
func (c *CSVBytesProcessor2) SetHasHeader(hasHeader bool) *CSVBytesProcessor2 {
	c.hasHeader = hasHeader
	return c
}

// SetDelimiter 设置分隔符
func (c *CSVBytesProcessor2) SetDelimiter(delimiter rune) *CSVBytesProcessor2 {
	c.delimiter = delimiter
	return c
}

// SetSkipRows 设置跳过的行数
func (c *CSVBytesProcessor2) SetSkipRows(skipRowsBegin, skipRowsEnd int) *CSVBytesProcessor2 {
	c.skipRowsBegin = skipRowsBegin
	c.skipRowsEnd = skipRowsEnd
	return c
}

// ReadFromBytes 从二进制流读取 CSV 数据到 struct 切片
// data 必须是指向切片的指针，例如: &[]User{}
func (c *CSVBytesProcessor2) ReadFromBytes(csvData []byte, data interface{}) error {
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

// ReadFromBytesWithSkip 从二进制流读取 CSV 数据到 struct 切片，支持跳过指定行数
// data 必须是指向切片的指针，例如: &[]User{}
func (c *CSVBytesProcessor2) ReadFromBytesWithSkip(csvData []byte, data interface{}) error {
	if len(csvData) == 0 {
		return fmt.Errorf("CSV 数据为空")
	}

	// 按行分割
	lines := bytes.Split(csvData, []byte("\n"))

	var beforeLines [][]byte
	var afterLines [][]byte

	// 只获取指定位置的行合并 skipRowsBegin 行之前的数据加上skipRowsEnd之后行的数据
	if c.skipRowsBegin > 0 && len(lines) > c.skipRowsBegin {
		beforeLines = lines[:c.skipRowsBegin-1]
	}
	if c.skipRowsEnd > 0 && len(lines) > c.skipRowsEnd {
		afterLines = lines[c.skipRowsEnd:]
	}
	lines = append(beforeLines, afterLines...)
	// 重新组合数据
	newData := bytes.Join(lines, []byte("\n"))

	// 创建字节读取器
	reader := bytes.NewReader(newData)

	// 使用 gocsv 从字节读取器读取数据到 struct
	err := gocsv.Unmarshal(reader, data)
	if err != nil {
		return fmt.Errorf("解析 CSV 二进制流失败: %v", err)
	}

	return nil
}

// WriteToBytes 将 struct 切片写入二进制流
// data 必须是切片，例如: []User{}
func (c *CSVBytesProcessor2) WriteToBytes(data interface{}) ([]byte, error) {
	// 创建字节缓冲区
	var buf bytes.Buffer

	// 使用 gocsv 将 struct 写入字节缓冲区
	err := gocsv.Marshal(data, &buf)
	if err != nil {
		return nil, fmt.Errorf("写入 CSV 到二进制流失败: %v", err)
	}

	return buf.Bytes(), nil
}

// GetDelimiter 获取当前设置的分隔符
func (c *CSVBytesProcessor2) GetDelimiter() rune {
	return c.delimiter
}

// GetHasHeader 获取是否包含标题行
func (c *CSVBytesProcessor2) GetHasHeader() bool {
	return c.hasHeader
}

// GetSkipRows 获取跳过的行数
func (c *CSVBytesProcessor2) GetSkipRows() (int, int) {
	return c.skipRowsBegin, c.skipRowsEnd
}
