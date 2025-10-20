package excel

import (
	"fmt"
	"os"
	"testing"

	"github.com/xuri/excelize/v2"
)

// TestExcelReader_ReadToStruct 测试读取Excel到结构体
func TestExcelReader_ReadToStruct(t *testing.T) {
	// 创建测试Excel文件
	testFile := "test_read.xlsx"
	createTestExcelFile(testFile)
	defer os.Remove(testFile)

	tests := []struct {
		name     string
		filePath string
		sheet    string
		wantErr  bool
	}{
		{
			name:     "正常读取用户数据",
			filePath: testFile,
			sheet:    "用户信息",
			wantErr:  false,
		},
		{
			name:     "读取不存在的文件",
			filePath: "not_exist.xlsx",
			sheet:    "Sheet1",
			wantErr:  true,
		},
		{
			name:     "读取不存在的sheet",
			filePath: testFile,
			sheet:    "不存在的Sheet",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var users []User
			err := NewExcelReader(tt.filePath).
				SetSheet(tt.sheet).
				ReadToStruct(&users)

			if (err != nil) != tt.wantErr {
				t.Errorf("ReadToStruct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if len(users) == 0 {
					t.Error("期望读取到数据，但结果为空")
					return
				}

				// 验证第一个用户的数据
				firstUser := users[0]
				if firstUser.ID != 1 {
					t.Errorf("期望ID=1，实际ID=%d", firstUser.ID)
				}
				if firstUser.Name != "张三" {
					t.Errorf("期望姓名=张三，实际姓名=%s", firstUser.Name)
				}
				if firstUser.Age != 25 {
					t.Errorf("期望年龄=25，实际年龄=%d", firstUser.Age)
				}
				if firstUser.Salary != 5000.50 {
					t.Errorf("期望薪资=5000.50，实际薪资=%f", firstUser.Salary)
				}
				if !firstUser.IsActive {
					t.Error("期望IsActive=true，实际为false")
				}

				t.Logf("成功读取 %d 条用户数据", len(users))
				for i, user := range users {
					t.Logf("用户%d: ID=%d, 姓名=%s, 年龄=%d, 薪资=%f", i+1, user.ID, user.Name, user.Age, user.Salary)
				}
			}
		})
	}
}

// TestExcelReader_ReadToMap 测试读取Excel到Map
func TestExcelReader_ReadToMap(t *testing.T) {
	// 创建测试Excel文件
	testFile := "test_read_map.xlsx"
	createTestExcelFile(testFile)
	defer os.Remove(testFile)

	var users []User
	err := NewExcelReader(testFile).
		SetSheet("用户信息").
		ReadToStruct(&users)

	if err != nil {
		t.Fatalf("读取失败: %v", err)
	}

	// 使用Map方式读取相同数据
	maps, err := NewExcelReader(testFile).
		SetSheet("用户信息").
		ReadToMap()

	if err != nil {
		t.Fatalf("Map读取失败: %v", err)
	}

	if len(maps) != len(users) {
		t.Errorf("期望读取%d条数据，实际读取%d条", len(users), len(maps))
	}

	t.Logf("Map方式读取到 %d 条数据", len(maps))
	for i, m := range maps {
		t.Logf("第%d条: %v", i+1, m)
	}
}

// TestExcelReader_CustomSettings 测试自定义设置
func TestExcelReader_CustomSettings(t *testing.T) {
	// 创建自定义格式的测试文件
	testFile := "test_custom.xlsx"
	createCustomTestExcelFile(testFile)
	defer os.Remove(testFile)

	var products []Product
	err := NewExcelReader(testFile).
		SetSheet("产品信息").
		SetHeaderRow(2).    // 表头在第2行
		SetDataStartRow(4). // 数据从第4行开始
		ReadToStruct(&products)

	if err != nil {
		t.Fatalf("读取失败: %v", err)
	}

	if len(products) == 0 {
		t.Error("期望读取到产品数据，但结果为空")
		return
	}

	t.Logf("成功读取 %d 条产品数据", len(products))
	for i, product := range products {
		t.Logf("产品%d: ID=%d, 名称=%s, 价格=%f", i+1, product.ID, product.Name, product.Price)
	}
}

// TestExcelReader_ComplexTypes 测试复杂类型解析
func TestExcelReader_ComplexTypes(t *testing.T) {
	// 创建包含复杂类型的测试文件
	testFile := "test_complex.xlsx"
	createComplexTestExcelFile(testFile)
	defer os.Remove(testFile)

	var employees []Employee
	err := NewExcelReader(testFile).
		SetSheet("员工信息").
		ReadToStruct(&employees)

	if err != nil {
		t.Fatalf("读取失败: %v", err)
	}

	if len(employees) == 0 {
		t.Error("期望读取到员工数据，但结果为空")
		return
	}

	t.Logf("成功读取 %d 条员工数据", len(employees))
	for i, emp := range employees {
		t.Logf("员工%d: ID=%d, 姓名=%s, 部门=%s, 入职日期=%s",
			i+1, emp.ID, emp.Name, emp.Department, emp.HireDate.Format("2006-01-02"))
	}
}

// createTestExcelFile 创建测试用的Excel文件
func createTestExcelFile(filename string) {
	f := excelize.NewFile()

	// 创建用户信息sheet
	sheetName := "用户信息"
	index, _ := f.NewSheet(sheetName)
	f.SetActiveSheet(index)

	// 设置表头
	headers := []string{"ID", "姓名", "年龄", "邮箱", "电话", "薪资", "是否激活", "生日", "创建时间"}
	for i, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheetName, cell, header)
	}

	// 设置测试数据
	testData := [][]interface{}{
		{1, "张三", 25, "zhangsan@example.com", "13800138001", 5000.50, true, "1998-05-15", "2023-01-01 09:00:00"},
		{2, "李四", 30, "lisi@example.com", "13800138002", 6000.00, true, "1993-08-20", "2023-01-02 10:30:00"},
		{3, "王五", 28, "wangwu@example.com", "13800138003", 5500.75, false, "1995-12-10", "2023-01-03 14:15:00"},
	}

	for rowIndex, row := range testData {
		for colIndex, value := range row {
			cell, _ := excelize.CoordinatesToCellName(colIndex+1, rowIndex+2)
			f.SetCellValue(sheetName, cell, value)
		}
	}

	f.SaveAs(filename)
}

// createCustomTestExcelFile 创建自定义格式的测试文件
func createCustomTestExcelFile(filename string) {
	f := excelize.NewFile()

	// 创建产品信息sheet
	sheetName := "产品信息"
	index, _ := f.NewSheet(sheetName)
	f.SetActiveSheet(index)

	// 第1行：标题
	f.SetCellValue(sheetName, "A1", "产品信息表")

	// 第2行：表头
	headers := []string{"产品ID", "产品名称", "分类", "价格", "库存", "描述", "是否可用"}
	for i, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 2)
		f.SetCellValue(sheetName, cell, header)
	}

	// 第3行：空行

	// 第4行开始：数据
	testData := [][]interface{}{
		{1, "iPhone 15", "手机", 6999.00, 100, "最新款iPhone", true},
		{2, "MacBook Pro", "电脑", 12999.00, 50, "专业级笔记本电脑", true},
		{3, "AirPods Pro", "耳机", 1999.00, 200, "降噪无线耳机", false},
	}

	for rowIndex, row := range testData {
		for colIndex, value := range row {
			cell, _ := excelize.CoordinatesToCellName(colIndex+1, rowIndex+4)
			f.SetCellValue(sheetName, cell, value)
		}
	}

	f.SaveAs(filename)
}

// createComplexTestExcelFile 创建包含复杂类型的测试文件
func createComplexTestExcelFile(filename string) {
	f := excelize.NewFile()

	// 创建员工信息sheet
	sheetName := "员工信息"
	index, _ := f.NewSheet(sheetName)
	f.SetActiveSheet(index)

	// 设置表头
	headers := []string{"员工ID", "姓名", "部门", "职位", "薪资", "入职日期", "是否经理", "电话号码", "邮箱地址"}
	for i, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheetName, cell, header)
	}

	// 设置测试数据
	testData := [][]interface{}{
		{1, "张经理", "技术部", "部门经理", 15000.00, "2020-01-15", true, "13800138001", "zhang@company.com"},
		{2, "李工程师", "技术部", "高级工程师", 12000.00, "2021-03-20", false, "13800138002", "li@company.com"},
		{3, "王设计师", "设计部", "UI设计师", 10000.00, "2022-06-10", false, "13800138003", "wang@company.com"},
	}

	for rowIndex, row := range testData {
		for colIndex, value := range row {
			cell, _ := excelize.CoordinatesToCellName(colIndex+1, rowIndex+2)
			f.SetCellValue(sheetName, cell, value)
		}
	}

	f.SaveAs(filename)
}

// TestExcelReader_ErrorHandling 测试错误处理
func TestExcelReader_ErrorHandling(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() string
		wantErr bool
	}{
		{
			name: "空文件",
			setup: func() string {
				f := excelize.NewFile()
				filename := "empty.xlsx"
				f.SaveAs(filename)
				return filename
			},
			wantErr: true,
		},
		{
			name: "只有表头没有数据",
			setup: func() string {
				f := excelize.NewFile()
				filename := "header_only.xlsx"
				f.SetCellValue("Sheet1", "A1", "ID")
				f.SetCellValue("Sheet1", "B1", "Name")
				f.SaveAs(filename)
				return filename
			},
			wantErr: false, // 应该成功，只是没有数据
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filename := tt.setup()
			defer os.Remove(filename)

			var users []User
			err := NewExcelReader(filename).ReadToStruct(&users)

			if (err != nil) != tt.wantErr {
				t.Errorf("ReadToStruct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// BenchmarkExcelReader_ReadToStruct 性能测试
func BenchmarkExcelReader_ReadToStruct(b *testing.B) {
	// 创建大量数据的测试文件
	testFile := "benchmark_test.xlsx"
	createLargeTestExcelFile(testFile, 1000) // 1000条数据
	defer os.Remove(testFile)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var users []User
		err := NewExcelReader(testFile).
			SetSheet("用户信息").
			ReadToStruct(&users)
		if err != nil {
			b.Fatalf("读取失败: %v", err)
		}
	}
}

// createLargeTestExcelFile 创建大量数据的测试文件
func createLargeTestExcelFile(filename string, count int) {
	f := excelize.NewFile()

	sheetName := "用户信息"
	index, _ := f.NewSheet(sheetName)
	f.SetActiveSheet(index)

	// 设置表头
	headers := []string{"ID", "姓名", "年龄", "邮箱", "电话", "薪资", "是否激活", "生日", "创建时间"}
	for i, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheetName, cell, header)
	}

	// 生成大量测试数据
	for i := 1; i <= count; i++ {
		row := []interface{}{
			i,
			fmt.Sprintf("用户%d", i),
			20 + (i % 40),
			fmt.Sprintf("user%d@example.com", i),
			fmt.Sprintf("138%08d", i),
			3000.0 + float64(i*10),
			i%2 == 0,
			fmt.Sprintf("1990-%02d-%02d", (i%12)+1, (i%28)+1),
			fmt.Sprintf("2023-01-%02d 09:00:00", (i%28)+1),
		}

		for colIndex, value := range row {
			cell, _ := excelize.CoordinatesToCellName(colIndex+1, i+1)
			f.SetCellValue(sheetName, cell, value)
		}
	}

	f.SaveAs(filename)
}
