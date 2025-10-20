package excel

import (
	"fmt"
	"reflect"
	"time"

	"github.com/spf13/cast"
	"github.com/xuri/excelize/v2"
)

// StructExcelWriter 将结构体切片导出为 Excel 的工具
type StructExcelWriter struct {
	data          reflect.Value
	sheetName     string
	path          string
	headerFromTag bool
	timeFormat    string
	columns       []string // 按字段名指定导出列顺序（可选），不指定则按结构体字段顺序
	headers       []string // 自定义表头（可选），不指定则使用 excel 标签或字段名
	Error         error
}

// NewStructExcelWriter 构造函数
func NewStructExcelWriter(slice interface{}) *StructExcelWriter {
	val := reflect.ValueOf(slice)
	if val.Kind() != reflect.Slice {
		return &StructExcelWriter{Error: fmt.Errorf("input is not a slice")}
	}
	return &StructExcelWriter{
		data:       val,
		sheetName:  "Sheet1",
		path:       "export.xlsx",
		timeFormat: "2006-01-02 15:04:05",
	}
}

// Sheet 指定工作表名称
func (w *StructExcelWriter) Sheet(sheet string) *StructExcelWriter {
	if w.Error != nil {
		return w
	}
	if sheet == "" {
		w.Error = fmt.Errorf("sheet name is empty")
		return w
	}
	w.sheetName = sheet
	return w
}

// SavePath 指定保存路径
func (w *StructExcelWriter) SavePath(path string) *StructExcelWriter {
	if w.Error != nil {
		return w
	}
	if path == "" {
		w.Error = fmt.Errorf("path is empty")
		return w
	}
	w.path = path
	return w
}

// UseTagHeaders 使用结构体 `excel` 标签作为表头
func (w *StructExcelWriter) UseTagHeaders() *StructExcelWriter {
	w.headerFromTag = true
	return w
}

// TimeFormat 指定 time.Time 格式化字符串
func (w *StructExcelWriter) TimeFormat(layout string) *StructExcelWriter {
	if layout != "" {
		w.timeFormat = layout
	}
	return w
}

// Columns 指定导出字段（通过字段名），并决定列顺序
func (w *StructExcelWriter) Columns(columns ...string) *StructExcelWriter {
	w.columns = append([]string{}, columns...)
	return w
}

// Headers 指定自定义表头（与 Columns 一一对应）
func (w *StructExcelWriter) Headers(headers ...string) *StructExcelWriter {
	w.headers = append([]string{}, headers...)
	return w
}

// ToExcel 执行导出
func (w *StructExcelWriter) ToExcel() *StructExcelWriter {
	if w.Error != nil {
		return w
	}
	// 空数据直接创建空表（仅表头）
	if w.data.Len() == 0 {
		f := excelize.NewFile()
		_ = f.SetSheetName("Sheet1", w.sheetName)
		if len(w.headers) > 0 {
			_ = f.SetSheetRow(w.sheetName, "A1", &w.headers)
		}
		w.Error = f.SaveAs(w.path)
		return w
	}

	// 获取元素类型（去掉指针）
	elemType := w.data.Type().Elem()
	if elemType.Kind() == reflect.Ptr {
		elemType = elemType.Elem()
	}
	if elemType.Kind() != reflect.Struct {
		w.Error = fmt.Errorf("slice element must be struct or *struct")
		return w
	}

	// 计算导出列
	fieldIndexByName := make(map[string]int)
	for i := 0; i < elemType.NumField(); i++ {
		field := elemType.Field(i)
		// 跳过不可导出的非导出字段
		if field.PkgPath != "" { // 非导出
			continue
		}
		fieldIndexByName[field.Name] = i
	}

	var exportFieldIndexes []int
	if len(w.columns) > 0 {
		for _, name := range w.columns {
			idx, ok := fieldIndexByName[name]
			if !ok {
				w.Error = fmt.Errorf("column %s not found in struct", name)
				return w
			}
			exportFieldIndexes = append(exportFieldIndexes, idx)
		}
	} else {
		// 默认按结构体字段顺序
		for i := 0; i < elemType.NumField(); i++ {
			field := elemType.Field(i)
			if field.PkgPath != "" {
				continue
			}
			exportFieldIndexes = append(exportFieldIndexes, i)
		}
	}

	// 生成表头
	headers := make([]string, 0, len(exportFieldIndexes))
	if len(w.headers) > 0 {
		headers = append(headers, w.headers...)
	} else {
		for _, idx := range exportFieldIndexes {
			field := elemType.Field(idx)
			if w.headerFromTag {
				if tag := field.Tag.Get("excel"); tag != "" {
					headers = append(headers, tag)
					continue
				}
			}
			headers = append(headers, field.Name)
		}
	}

	// 写入 Excel
	f := excelize.NewFile()
	_ = f.SetSheetName("Sheet1", w.sheetName)
	// 表头
	_ = f.SetSheetRow(w.sheetName, "A1", &headers)

	// 数据
	for i := 0; i < w.data.Len(); i++ {
		rowValues := make([]interface{}, 0, len(exportFieldIndexes))
		rowVal := w.data.Index(i)
		if rowVal.Kind() == reflect.Ptr {
			if rowVal.IsNil() {
				rowValues = append(rowValues, make([]interface{}, len(exportFieldIndexes))...)
				continue
			}
			rowVal = rowVal.Elem()
		}
		for _, idx := range exportFieldIndexes {
			fieldVal := rowVal.Field(idx)
			rowValues = append(rowValues, w.stringify(fieldVal))
		}
		cell, _ := excelize.CoordinatesToCellName(1, i+2)
		_ = f.SetSheetRow(w.sheetName, cell, &rowValues)
	}

	w.Error = f.SaveAs(w.path)
	return w
}

// stringify 将任意字段值转换为适合写入 Excel 的值
func (w *StructExcelWriter) stringify(v reflect.Value) interface{} {
	if !v.IsValid() {
		return ""
	}
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return ""
		}
		return w.stringify(v.Elem())
	}
	// time.Time 特殊处理
	if v.Type() == reflect.TypeOf(time.Time{}) {
		return v.Interface().(time.Time).Format(w.timeFormat)
	}
	switch v.Kind() {
	case reflect.String:
		return v.String()
	case reflect.Bool:
		return v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint()
	case reflect.Float32, reflect.Float64:
		return v.Float()
	case reflect.Struct, reflect.Map, reflect.Slice, reflect.Array:
		return cast.ToString(v.Interface())
	default:
		return cast.ToString(v.Interface())
	}
}
