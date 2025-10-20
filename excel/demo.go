package excel

import (
	"fmt"
	"os"

	"github.com/xuri/excelize/v2"
)

// CreateSampleExcelFile 创建示例Excel文件用于测试
func CreateSampleExcelFile(filename string) error {
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

	// 设置示例数据
	sampleData := [][]interface{}{
		{1, "张三", 25, "zhangsan@example.com", "13800138001", 5000.50, true, "1998-05-15", "2023-01-01 09:00:00"},
		{2, "李四", 30, "lisi@example.com", "13800138002", 6000.00, true, "1993-08-20", "2023-01-02 10:30:00"},
		{3, "王五", 28, "wangwu@example.com", "13800138003", 5500.75, false, "1995-12-10", "2023-01-03 14:15:00"},
		{4, "赵六", 35, "zhaoliu@example.com", "13800138004", 7000.00, true, "1988-03-25", "2023-01-04 16:45:00"},
		{5, "钱七", 22, "qianqi@example.com", "13800138005", 4500.00, true, "2001-07-08", "2023-01-05 11:20:00"},
	}

	for rowIndex, row := range sampleData {
		for colIndex, value := range row {
			cell, _ := excelize.CoordinatesToCellName(colIndex+1, rowIndex+2)
			f.SetCellValue(sheetName, cell, value)
		}
	}

	// 创建产品信息sheet
	productSheet := "产品信息"
	f.NewSheet(productSheet)

	// 设置产品表头
	productHeaders := []string{"产品ID", "产品名称", "分类", "价格", "库存", "描述", "是否可用"}
	for i, header := range productHeaders {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(productSheet, cell, header)
	}

	// 设置产品示例数据
	productData := [][]interface{}{
		{1, "iPhone 15", "手机", 6999.00, 100, "最新款iPhone", true},
		{2, "MacBook Pro", "电脑", 12999.00, 50, "专业级笔记本电脑", true},
		{3, "AirPods Pro", "耳机", 1999.00, 200, "降噪无线耳机", true},
		{4, "iPad Air", "平板", 4399.00, 80, "轻薄平板电脑", true},
		{5, "Apple Watch", "手表", 2999.00, 150, "智能手表", false},
	}

	for rowIndex, row := range productData {
		for colIndex, value := range row {
			cell, _ := excelize.CoordinatesToCellName(colIndex+1, rowIndex+2)
			f.SetCellValue(productSheet, cell, value)
		}
	}

	// 创建员工信息sheet
	employeeSheet := "员工信息"
	f.NewSheet(employeeSheet)

	// 设置员工表头
	employeeHeaders := []string{"员工ID", "姓名", "部门", "职位", "薪资", "入职日期", "是否经理", "电话号码", "邮箱地址"}
	for i, header := range employeeHeaders {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(employeeSheet, cell, header)
	}

	// 设置员工示例数据
	employeeData := [][]interface{}{
		{1, "张经理", "技术部", "部门经理", 15000.00, "2020-01-15", true, "13800138001", "zhang@company.com"},
		{2, "李工程师", "技术部", "高级工程师", 12000.00, "2021-03-20", false, "13800138002", "li@company.com"},
		{3, "王设计师", "设计部", "UI设计师", 10000.00, "2022-06-10", false, "13800138003", "wang@company.com"},
		{4, "刘产品", "产品部", "产品经理", 13000.00, "2020-09-05", true, "13800138004", "liu@company.com"},
		{5, "陈运营", "运营部", "运营专员", 8000.00, "2023-02-14", false, "13800138005", "chen@company.com"},
	}

	for rowIndex, row := range employeeData {
		for colIndex, value := range row {
			cell, _ := excelize.CoordinatesToCellName(colIndex+1, rowIndex+2)
			f.SetCellValue(employeeSheet, cell, value)
		}
	}

	// 保存文件
	err := f.SaveAs(filename)
	if err != nil {
		return fmt.Errorf("保存文件失败: %v", err)
	}

	fmt.Printf("示例Excel文件已创建: %s\n", filename)
	fmt.Println("包含以下工作表:")
	fmt.Println("- 用户信息: 5条用户数据")
	fmt.Println("- 产品信息: 5条产品数据")
	fmt.Println("- 员工信息: 5条员工数据")

	return nil
}

// CreateSampleExcelFileWithCustomFormat 创建自定义格式的示例Excel文件
func CreateSampleExcelFileWithCustomFormat(filename string) error {
	f := excelize.NewFile()

	// 创建自定义格式的产品信息sheet
	sheetName := "产品信息"
	index, _ := f.NewSheet(sheetName)
	f.SetActiveSheet(index)

	// 第1行：标题
	f.SetCellValue(sheetName, "A1", "产品信息表")
	f.MergeCell(sheetName, "A1", "G1")

	// 第2行：表头
	headers := []string{"产品ID", "产品名称", "分类", "价格", "库存", "描述", "是否可用"}
	for i, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 2)
		f.SetCellValue(sheetName, cell, header)
	}

	// 第3行：空行

	// 第4行开始：数据
	sampleData := [][]interface{}{
		{1, "iPhone 15", "手机", 6999.00, 100, "最新款iPhone", true},
		{2, "MacBook Pro", "电脑", 12999.00, 50, "专业级笔记本电脑", true},
		{3, "AirPods Pro", "耳机", 1999.00, 200, "降噪无线耳机", true},
		{4, "iPad Air", "平板", 4399.00, 80, "轻薄平板电脑", true},
		{5, "Apple Watch", "手表", 2999.00, 150, "智能手表", false},
	}

	for rowIndex, row := range sampleData {
		for colIndex, value := range row {
			cell, _ := excelize.CoordinatesToCellName(colIndex+1, rowIndex+4)
			f.SetCellValue(sheetName, cell, value)
		}
	}

	// 保存文件
	err := f.SaveAs(filename)
	if err != nil {
		return fmt.Errorf("保存文件失败: %v", err)
	}

	fmt.Printf("自定义格式示例Excel文件已创建: %s\n", filename)
	fmt.Println("表头在第2行，数据从第4行开始")

	return nil
}

// DemoExcelReader 演示Excel读取功能
func DemoExcelReader() {
	// 创建示例文件
	sampleFile := "sample_data.xlsx"
	err := CreateSampleExcelFile(sampleFile)
	if err != nil {
		fmt.Printf("创建示例文件失败: %v\n", err)
		return
	}
	defer os.Remove(sampleFile)

	fmt.Println("\n=== Excel读取功能演示 ===")

	// 演示1: 读取用户信息
	fmt.Println("\n1. 读取用户信息:")
	var users []User
	err = NewExcelReader(sampleFile).
		SetSheet("用户信息").
		ReadToStruct(&users)

	if err != nil {
		fmt.Printf("读取用户信息失败: %v\n", err)
	} else {
		fmt.Printf("成功读取 %d 条用户数据:\n", len(users))
		for i, user := range users {
			fmt.Printf("  用户%d: ID=%d, 姓名=%s, 年龄=%d, 薪资=%.2f, 激活=%t\n",
				i+1, user.ID, user.Name, user.Age, user.Salary, user.IsActive)
		}
	}

	// 演示2: 读取产品信息
	fmt.Println("\n2. 读取产品信息:")
	var products []Product
	err = NewExcelReader(sampleFile).
		SetSheet("产品信息").
		ReadToStruct(&products)

	if err != nil {
		fmt.Printf("读取产品信息失败: %v\n", err)
	} else {
		fmt.Printf("成功读取 %d 条产品数据:\n", len(products))
		for i, product := range products {
			fmt.Printf("  产品%d: ID=%d, 名称=%s, 分类=%s, 价格=%.2f, 库存=%d\n",
				i+1, product.ID, product.Name, product.Category, product.Price, product.Stock)
		}
	}

	// 演示3: 读取员工信息
	fmt.Println("\n3. 读取员工信息:")
	var employees []Employee
	err = NewExcelReader(sampleFile).
		SetSheet("员工信息").
		ReadToStruct(&employees)

	if err != nil {
		fmt.Printf("读取员工信息失败: %v\n", err)
	} else {
		fmt.Printf("成功读取 %d 条员工数据:\n", len(employees))
		for i, emp := range employees {
			fmt.Printf("  员工%d: ID=%d, 姓名=%s, 部门=%s, 职位=%s, 薪资=%.2f\n",
				i+1, emp.ID, emp.Name, emp.Department, emp.Position, emp.Salary)
		}
	}

	// 演示4: 使用Map方式读取
	fmt.Println("\n4. 使用Map方式读取:")
	maps, err := NewExcelReader(sampleFile).
		SetSheet("用户信息").
		ReadToMap()

	if err != nil {
		fmt.Printf("Map方式读取失败: %v\n", err)
	} else {
		fmt.Printf("Map方式成功读取 %d 条数据:\n", len(maps))
		for i, m := range maps {
			fmt.Printf("  第%d条: %v\n", i+1, m)
		}
	}

	fmt.Println("\n=== 演示完成 ===")
}
