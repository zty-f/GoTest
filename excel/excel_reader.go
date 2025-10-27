package excel

import (
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"
)

// ExcelReader Excel读取器
type ExcelReader struct {
	filePath     string
	reader       io.ReadSeeker
	url          string
	sheetName    string
	headerRow    int // 表头行号，从1开始
	dataStartRow int // 数据开始行号，从1开始
	error        error
}

// NewExcelReader 创建新的Excel读取器（从文件路径）
func NewExcelReader(filePath string) *ExcelReader {
	return &ExcelReader{
		filePath:     filePath,
		sheetName:    "Sheet1",
		headerRow:    1,
		dataStartRow: 2,
	}
}

// NewExcelReaderFromReader 从io.ReadSeeker创建Excel读取器（文件流）
func NewExcelReaderFromReader(reader io.ReadSeeker) *ExcelReader {
	return &ExcelReader{
		reader:       reader,
		sheetName:    "Sheet1",
		headerRow:    1,
		dataStartRow: 2,
	}
}

// NewExcelReaderFromURL 从URL创建Excel读取器
func NewExcelReaderFromURL(url string) *ExcelReader {
	return &ExcelReader{
		url:          url,
		sheetName:    "Sheet1",
		headerRow:    1,
		dataStartRow: 2,
	}
}

// SetSheet 设置工作表名称
func (r *ExcelReader) SetSheet(sheetName string) *ExcelReader {
	if r.error != nil {
		return r
	}
	r.sheetName = sheetName
	return r
}

// SetHeaderRow 设置表头行号
func (r *ExcelReader) SetHeaderRow(row int) *ExcelReader {
	if r.error != nil {
		return r
	}
	if row < 1 {
		r.error = fmt.Errorf("header row must be >= 1")
		return r
	}
	r.headerRow = row
	return r
}

// SetDataStartRow 设置数据开始行号
func (r *ExcelReader) SetDataStartRow(row int) *ExcelReader {
	if r.error != nil {
		return r
	}
	if row < 1 {
		r.error = fmt.Errorf("data start row must be >= 1")
		return r
	}
	r.dataStartRow = row
	return r
}

// openExcelFile 打开Excel文件，支持文件路径、文件流和URL三种方式
func (r *ExcelReader) openExcelFile() (*excelize.File, error) {
	if r.error != nil {
		return nil, r.error
	}

	// 优先使用文件流
	if r.reader != nil {
		f, err := excelize.OpenReader(r.reader)
		if err != nil {
			return nil, fmt.Errorf("failed to open reader: %v", err)
		}
		return f, nil
	}

	// 其次使用URL
	if r.url != "" {
		resp, err := http.Get(r.url)
		if err != nil {
			return nil, fmt.Errorf("failed to download from URL: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("HTTP error: %s", resp.Status)
		}

		f, err := excelize.OpenReader(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to open from URL: %v", err)
		}
		return f, nil
	}

	// 最后使用文件路径
	if r.filePath != "" {
		f, err := excelize.OpenFile(r.filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to open file: %v", err)
		}
		return f, nil
	}

	return nil, fmt.Errorf("no valid data source provided (file path, reader, or URL)")
}

// ReadToStruct 读取Excel数据到结构体切片
func (r *ExcelReader) ReadToStruct(result interface{}) error {
	if r.error != nil {
		return r.error
	}

	// 检查result是否为切片指针
	resultValue := reflect.ValueOf(result)
	if resultValue.Kind() != reflect.Ptr {
		return fmt.Errorf("result must be a pointer to slice")
	}

	resultElem := resultValue.Elem()
	if resultElem.Kind() != reflect.Slice {
		return fmt.Errorf("result must be a pointer to slice")
	}

	// 打开Excel文件（支持文件路径、文件流和URL）
	f, err := r.openExcelFile()
	if err != nil {
		return err
	}
	defer f.Close()

	// 获取所有行
	rows, err := f.GetRows(r.sheetName)
	if err != nil {
		return fmt.Errorf("failed to get rows: %v", err)
	}

	if len(rows) == 0 {
		return fmt.Errorf("no data found in sheet %s", r.sheetName)
	}

	// 检查表头行是否存在
	if r.headerRow > len(rows) {
		return fmt.Errorf("header row %d exceeds total rows %d", r.headerRow, len(rows))
	}

	// 获取表头
	headers := rows[r.headerRow-1]
	if len(headers) == 0 {
		return fmt.Errorf("no headers found in row %d", r.headerRow)
	}

	// 获取结构体类型
	elemType := resultElem.Type().Elem()
	if elemType.Kind() == reflect.Ptr {
		elemType = elemType.Elem()
	}

	// 创建字段映射
	fieldMap := make(map[string]int)
	for i := 0; i < elemType.NumField(); i++ {
		field := elemType.Field(i)
		tag := field.Tag.Get("excel")
		if tag != "" {
			fieldMap[tag] = i
		} else {
			// 如果没有excel标签，使用字段名
			fieldMap[field.Name] = i
		}
	}

	// 创建列索引映射
	columnMap := make(map[int]int) // 列索引 -> 字段索引
	for colIndex, header := range headers {
		if fieldIndex, exists := fieldMap[header]; exists {
			columnMap[colIndex] = fieldIndex
		}
	}

	// 读取数据行
	for rowIndex := r.dataStartRow - 1; rowIndex < len(rows); rowIndex++ {
		row := rows[rowIndex]
		if len(row) == 0 {
			continue // 跳过空行
		}

		// 创建结构体实例
		var elem reflect.Value
		if resultElem.Type().Elem().Kind() == reflect.Ptr {
			elem = reflect.New(elemType)
		} else {
			elem = reflect.New(elemType).Elem()
		}

		// 填充字段值
		for colIndex, fieldIndex := range columnMap {
			if colIndex >= len(row) {
				continue
			}

			cellValue := row[colIndex]
			if cellValue == "" {
				continue
			}

			field := elemType.Field(fieldIndex)
			fieldValue := elem
			if elem.Kind() == reflect.Ptr {
				fieldValue = elem.Elem()
			}

			if err := r.setFieldValue(fieldValue.Field(fieldIndex), cellValue, field.Type); err != nil {
				return fmt.Errorf("failed to set field %s at row %d: %v", field.Name, rowIndex+1, err)
			}
		}

		// 添加到结果切片
		if resultElem.Type().Elem().Kind() == reflect.Ptr {
			resultElem.Set(reflect.Append(resultElem, elem))
		} else {
			resultElem.Set(reflect.Append(resultElem, elem))
		}
	}

	return nil
}

// setFieldValue 设置字段值
func (r *ExcelReader) setFieldValue(field reflect.Value, value string, fieldType reflect.Type) error {
	if !field.CanSet() {
		return fmt.Errorf("field cannot be set")
	}

	// 去除空格
	value = strings.TrimSpace(value)
	if value == "" {
		return nil
	}

	switch fieldType.Kind() {
	case reflect.String:
		field.SetString(value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if fieldType == reflect.TypeOf(time.Duration(0)) {
			// 特殊处理时间间隔
			duration, err := time.ParseDuration(value)
			if err != nil {
				return err
			}
			field.SetInt(int64(duration))
		} else {
			intVal, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return err
			}
			field.SetInt(intVal)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		uintVal, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return err
		}
		field.SetUint(uintVal)
	case reflect.Float32, reflect.Float64:
		floatVal, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		field.SetFloat(floatVal)
	case reflect.Bool:
		boolVal, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		field.SetBool(boolVal)
	case reflect.Struct:
		if fieldType == reflect.TypeOf(time.Time{}) {
			// 特殊处理时间类型
			timeVal, err := r.parseTime(value)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(timeVal))
		} else {
			return fmt.Errorf("unsupported struct type: %v", fieldType)
		}
	case reflect.Ptr:
		if field.IsNil() {
			field.Set(reflect.New(fieldType.Elem()))
		}
		return r.setFieldValue(field.Elem(), value, fieldType.Elem())
	default:
		return fmt.Errorf("unsupported field type: %v", fieldType)
	}

	return nil
}

// parseTime 解析时间字符串
func (r *ExcelReader) parseTime(value string) (time.Time, error) {
	// 尝试多种时间格式
	formats := []string{
		"2006-01-02 15:04:05",
		"2006-01-02",
		"2006/01/02 15:04:05",
		"2006/01/02",
		"01/02/2006 15:04:05",
		"01/02/2006",
		"2006-01-02T15:04:05Z",
		"2006-01-02T15:04:05.000Z",
		time.RFC3339,
		time.RFC3339Nano,
	}

	for _, format := range formats {
		if t, err := time.Parse(format, value); err == nil {
			return t, nil
		}
	}

	// 尝试解析时间戳
	if timestamp, err := strconv.ParseInt(value, 10, 64); err == nil {
		// 判断是秒还是毫秒时间戳
		if timestamp > 1e10 {
			return time.Unix(timestamp/1000, (timestamp%1000)*1e6), nil
		}
		return time.Unix(timestamp, 0), nil
	}

	return time.Time{}, fmt.Errorf("unable to parse time: %s", value)
}

// ReadToMap 读取Excel数据到map切片（备用方法）
func (r *ExcelReader) ReadToMap() ([]map[string]string, error) {
	if r.error != nil {
		return nil, r.error
	}

	// 打开Excel文件（支持文件路径、文件流和URL）
	f, err := r.openExcelFile()
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// 获取所有行
	rows, err := f.GetRows(r.sheetName)
	if err != nil {
		return nil, fmt.Errorf("failed to get rows: %v", err)
	}

	if len(rows) == 0 {
		return nil, fmt.Errorf("no data found in sheet %s", r.sheetName)
	}

	// 检查表头行是否存在
	if r.headerRow > len(rows) {
		return nil, fmt.Errorf("header row %d exceeds total rows %d", r.headerRow, len(rows))
	}

	// 获取表头
	headers := rows[r.headerRow-1]
	if len(headers) == 0 {
		return nil, fmt.Errorf("no headers found in row %d", r.headerRow)
	}

	var result []map[string]string

	// 读取数据行
	for rowIndex := r.dataStartRow - 1; rowIndex < len(rows); rowIndex++ {
		row := rows[rowIndex]
		if len(row) == 0 {
			continue // 跳过空行
		}

		rowMap := make(map[string]string)
		for colIndex, header := range headers {
			if colIndex < len(row) {
				rowMap[header] = row[colIndex]
			} else {
				rowMap[header] = ""
			}
		}

		result = append(result, rowMap)
	}

	return result, nil
}
