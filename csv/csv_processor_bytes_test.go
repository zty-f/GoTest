package csv

import (
	"fmt"
	"testing"
	"time"
)

// TestCSVBytesProcessor 测试 CSV 二进制流处理器
func TestCSVBytesProcessor(t *testing.T) {
	// 创建二进制流处理器
	processor := NewCSVBytesProcessor()

	// 设置配置
	processor.SetHasHeader(true).SetDelimiter(',')

	// 测试数据
	users := []User{
		{
			ID:        1,
			Name:      "张三",
			Age:       25,
			Email:     "zhangsan@example.com",
			City:      "北京",
			CreatedAt: time.Now(),
			IsActive:  true,
		},
		{
			ID:        2,
			Name:      "李四",
			Age:       30,
			Email:     "lisi@example.com",
			City:      "上海",
			CreatedAt: time.Now(),
			IsActive:  true,
		},
	}

	// 测试将结构体写入二进制流
	t.Run("WriteToBytes", func(t *testing.T) {
		csvBytes, err := processor.WriteToBytes(users)
		if err != nil {
			t.Fatalf("写入二进制流失败: %v", err)
		}

		if len(csvBytes) == 0 {
			t.Fatal("生成的二进制流为空")
		}

		fmt.Printf("成功生成 CSV 二进制流，大小: %d 字节\n", len(csvBytes))
		fmt.Printf("CSV 内容预览: %s\n", string(csvBytes[:min(100, len(csvBytes))]))
	})

	// 测试从二进制流读取到结构体
	t.Run("ReadFromBytes", func(t *testing.T) {
		// 先生成二进制流
		csvBytes, err := processor.WriteToBytes(users)
		if err != nil {
			t.Fatalf("生成二进制流失败: %v", err)
		}

		// 从二进制流读取
		var readUsers []User
		err = processor.ReadFromBytes(csvBytes, &readUsers)
		if err != nil {
			t.Fatalf("从二进制流读取失败: %v", err)
		}

		if len(readUsers) != len(users) {
			t.Fatalf("读取的数据数量不匹配，期望: %d，实际: %d", len(users), len(readUsers))
		}

		// 验证数据
		for i, user := range users {
			if readUsers[i].ID != user.ID || readUsers[i].Name != user.Name {
				t.Fatalf("数据不匹配，索引 %d，期望: %+v，实际: %+v", i, user, readUsers[i])
			}
		}

		fmt.Printf("成功从二进制流读取 %d 条用户数据\n", len(readUsers))
	})

	// 测试 Map 操作
	t.Run("MapOperations", func(t *testing.T) {
		// 测试数据
		mapData := []map[string]string{
			{"id": "1", "name": "张三", "age": "25"},
			{"id": "2", "name": "李四", "age": "30"},
		}

		// 写入二进制流
		csvBytes, err := processor.WriteMapToBytes(mapData)
		if err != nil {
			t.Fatalf("写入 Map 到二进制流失败: %v", err)
		}

		// 从二进制流读取
		readMapData, err := processor.ReadMapFromBytes(csvBytes)
		if err != nil {
			t.Fatalf("从二进制流读取 Map 失败: %v", err)
		}

		if len(readMapData) != len(mapData) {
			t.Fatalf("Map 数据数量不匹配")
		}

		fmt.Printf("成功处理 Map 数据，数量: %d\n", len(readMapData))
	})

	// 测试字符串操作
	t.Run("StringOperations", func(t *testing.T) {
		// 测试数据
		stringData := [][]string{
			{"id", "name", "age"},
			{"1", "张三", "25"},
			{"2", "李四", "30"},
		}

		// 写入二进制流
		csvBytes, err := processor.WriteStringToBytes(stringData)
		if err != nil {
			t.Fatalf("写入字符串到二进制流失败: %v", err)
		}

		// 从二进制流读取
		readStringData, err := processor.ReadStringFromBytes(csvBytes)
		if err != nil {
			t.Fatalf("从二进制流读取字符串失败: %v", err)
		}

		if len(readStringData) != len(stringData) {
			t.Fatalf("字符串数据数量不匹配")
		}

		fmt.Printf("成功处理字符串数据，行数: %d\n", len(readStringData))
	})

	// 测试格式验证
	t.Run("FormatValidation", func(t *testing.T) {
		// 生成有效的 CSV 二进制流
		csvBytes, err := processor.WriteToBytes(users)
		if err != nil {
			t.Fatalf("生成二进制流失败: %v", err)
		}

		// 验证格式
		err = processor.ValidateCSVFormat(csvBytes)
		if err != nil {
			t.Fatalf("格式验证失败: %v", err)
		}

		// 测试无效数据
		invalidData := []byte("这不是CSV数据")
		err = processor.ValidateCSVFormat(invalidData)
		if err == nil {
			t.Fatal("应该检测到无效的 CSV 格式")
		}

		fmt.Println("格式验证测试通过")
	})

	// 测试获取 CSV 信息
	t.Run("GetCSVInfo", func(t *testing.T) {
		csvBytes, err := processor.WriteToBytes(users)
		if err != nil {
			t.Fatalf("生成二进制流失败: %v", err)
		}

		info, err := processor.GetCSVInfo(csvBytes)
		if err != nil {
			t.Fatalf("获取 CSV 信息失败: %v", err)
		}

		fmt.Printf("CSV 信息: %+v\n", info)
	})

	// 测试配置获取
	t.Run("Configuration", func(t *testing.T) {
		delimiter := processor.GetDelimiter()
		hasHeader := processor.GetHasHeader()

		if delimiter != ',' {
			t.Fatalf("分隔符不匹配，期望: ','，实际: %c", delimiter)
		}

		if !hasHeader {
			t.Fatal("标题行设置不匹配")
		}

		fmt.Printf("配置验证通过 - 分隔符: %c, 包含标题行: %t\n", delimiter, hasHeader)
	})
}

// 辅助函数
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// TestCSVBytesProcessorWithRealData 测试使用真实 CSV 数据
func TestCSVBytesProcessorWithRealData(t *testing.T) {
	processor := NewCSVBytesProcessor()

	// 模拟前端上传的 CSV 二进制流
	csvData := []byte(`id,name,age,email,city,created_at,is_active
1,张三,25,zhangsan@example.com,北京,2024-01-01T00:00:00Z,true
2,李四,30,lisi@example.com,上海,2024-01-02T00:00:00Z,true
3,王五,28,wangwu@example.com,广州,2024-01-03T00:00:00Z,false`)

	t.Run("ReadRealCSVData", func(t *testing.T) {
		// 验证格式
		err := processor.ValidateCSVFormat(csvData)
		if err != nil {
			t.Fatalf("真实 CSV 数据格式验证失败: %v", err)
		}

		// 获取信息
		info, err := processor.GetCSVInfo(csvData)
		if err != nil {
			t.Fatalf("获取真实 CSV 信息失败: %v", err)
		}

		fmt.Printf("真实 CSV 数据信息: %+v\n", info)

		// 读取到结构体
		var users []User
		err = processor.ReadFromBytes(csvData, &users)
		if err != nil {
			t.Fatalf("读取真实 CSV 数据失败: %v", err)
		}

		if len(users) != 3 {
			t.Fatalf("读取的用户数量不匹配，期望: 3，实际: %d", len(users))
		}

		fmt.Printf("成功读取真实 CSV 数据，用户数量: %d\n", len(users))
		for i, user := range users {
			fmt.Printf("用户 %d: ID=%d, 姓名=%s, 年龄=%d, 城市=%s\n",
				i+1, user.ID, user.Name, user.Age, user.City)
		}
	})
}

// TestCSVBytesProcessorEdgeCases 测试边界情况
func TestCSVBytesProcessorEdgeCases(t *testing.T) {
	processor := NewCSVBytesProcessor()

	t.Run("EmptyData", func(t *testing.T) {
		// 测试空数据
		err := processor.ValidateCSVFormat([]byte{})
		if err == nil {
			t.Fatal("应该检测到空数据")
		}

		err = processor.ReadFromBytes([]byte{}, &[]User{})
		if err == nil {
			t.Fatal("应该检测到空数据")
		}
	})

	t.Run("SingleLineData", func(t *testing.T) {
		// 测试单行数据（没有换行符）
		singleLineData := []byte("id,name,age")
		err := processor.ValidateCSVFormat(singleLineData)
		if err == nil {
			t.Fatal("应该检测到缺少换行符")
		}
	})

	t.Run("NoDelimiter", func(t *testing.T) {
		// 测试没有分隔符的数据
		noDelimiterData := []byte("id\nname\nage")
		err := processor.ValidateCSVFormat(noDelimiterData)
		if err == nil {
			t.Fatal("应该检测到缺少分隔符")
		}
	})

	fmt.Println("边界情况测试通过")
}
