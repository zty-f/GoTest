package csv

import (
	"fmt"
	"testing"
	"time"
)

// User 用户结构体
type User struct {
	ID        int       `csv:"id"`         // CSV 列名映射
	Name      string    `csv:"name"`       // CSV 列名映射
	Age       int       `csv:"age"`        // CSV 列名映射
	Email     string    `csv:"email"`      // CSV 列名映射
	City      string    `csv:"city"`       // CSV 列名映射
	CreatedAt time.Time `csv:"created_at"` // CSV 列名映射
	IsActive  bool      `csv:"is_active"`  // CSV 列名映射
}

// Product 产品结构体
type Product struct {
	ID          int     `csv:"product_id"`
	Name        string  `csv:"product_name"`
	Price       float64 `csv:"price"`
	Category    string  `csv:"category"`
	Description string  `csv:"description"`
}

// TestWriteFromStruct 测试将结构体写入 CSV 文件
func TestWriteFromStruct(t *testing.T) {
	processor := NewCSVProcessor("test_users.csv")
	defer processor.DeleteFile() // 测试完成后删除文件

	// 创建测试数据
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
		{
			ID:        3,
			Name:      "王五",
			Age:       28,
			Email:     "wangwu@example.com",
			City:      "广州",
			CreatedAt: time.Now(),
			IsActive:  false,
		},
	}

	// 将结构体写入 CSV 文件
	err := processor.WriteFromStruct(users)
	if err != nil {
		t.Fatalf("写入 CSV 失败: %v", err)
	}

	// 检查文件是否创建成功
	if !processor.FileExists() {
		t.Fatal("CSV 文件未创建")
	}

	// 获取文件大小
	size, err := processor.GetFileSize()
	if err != nil {
		t.Fatalf("获取文件大小失败: %v", err)
	}
	if size == 0 {
		t.Fatal("CSV 文件为空")
	}

	fmt.Printf("成功写入 %d 条用户数据到 %s\n", len(users), processor.GetFilePath())
}

// TestReadToStruct 测试从 CSV 文件读取到结构体
func TestReadToStruct(t *testing.T) {
	processor := NewCSVProcessor("test_users.csv")

	// 先创建测试数据
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

	// 写入 CSV 文件
	err := processor.WriteFromStruct(users)
	if err != nil {
		t.Fatalf("写入 CSV 失败: %v", err)
	}

	// 从 CSV 文件读取到结构体
	var readUsers []User
	err = processor.ReadToStruct(&readUsers)
	if err != nil {
		t.Fatalf("读取 CSV 失败: %v", err)
	}

	// 验证数据
	if len(readUsers) != len(users) {
		t.Fatalf("读取的数据条数不匹配，期望 %d，实际 %d", len(users), len(readUsers))
	}

	// 验证第一条数据
	if readUsers[0].Name != users[0].Name {
		t.Fatalf("姓名不匹配，期望 %s，实际 %s", users[0].Name, readUsers[0].Name)
	}

	if readUsers[0].Age != users[0].Age {
		t.Fatalf("年龄不匹配，期望 %d，实际 %d", users[0].Age, readUsers[0].Age)
	}

	fmt.Printf("成功从 %s 读取 %d 条用户数据\n", processor.GetFilePath(), len(readUsers))
}

// TestReadToMap 测试读取 CSV 到 map
func TestReadToMap(t *testing.T) {
	processor := NewCSVProcessor("test_map.csv")
	defer processor.DeleteFile()

	// 创建 map 数据
	mapData := []map[string]string{
		{
			"id":    "1",
			"name":  "产品A",
			"price": "99.99",
		},
		{
			"id":    "2",
			"name":  "产品B",
			"price": "199.99",
		},
	}

	// 写入 CSV 文件
	err := processor.WriteFromMap(mapData)
	if err != nil {
		t.Fatalf("写入 CSV 失败: %v", err)
	}

	// 读取到 map
	readMapData, err := processor.ReadToMap()
	if err != nil {
		t.Fatalf("读取 CSV 失败: %v", err)
	}

	// 验证数据
	if len(readMapData) != len(mapData) {
		t.Fatalf("读取的数据条数不匹配，期望 %d，实际 %d", len(mapData), len(readMapData))
	}

	fmt.Printf("成功从 %s 读取 %d 条 map 数据\n", processor.GetFilePath(), len(readMapData))
}

// TestAppendToStruct 测试追加数据到现有 CSV 文件
func TestAppendToStruct(t *testing.T) {
	processor := NewCSVProcessor("test_append.csv")
	defer processor.DeleteFile()

	// 第一批数据
	users1 := []User{
		{
			ID:    1,
			Name:  "张三",
			Age:   25,
			Email: "zhangsan@example.com",
			City:  "北京",
		},
	}

	// 写入第一批数据
	err := processor.WriteFromStruct(users1)
	if err != nil {
		t.Fatalf("写入第一批数据失败: %v", err)
	}

	// 第二批数据
	users2 := []User{
		{
			ID:    2,
			Name:  "李四",
			Age:   30,
			Email: "lisi@example.com",
			City:  "上海",
		},
		{
			ID:    3,
			Name:  "王五",
			Age:   28,
			Email: "wangwu@example.com",
			City:  "广州",
		},
	}

	// 追加第二批数据
	err = processor.AppendToStruct(users2)
	if err != nil {
		t.Fatalf("追加数据失败: %v", err)
	}

	// 读取所有数据验证
	var allUsers []User
	err = processor.ReadToStruct(&allUsers)
	if err != nil {
		t.Fatalf("读取所有数据失败: %v", err)
	}

	// 验证总条数
	expectedCount := len(users1) + len(users2)
	if len(allUsers) != expectedCount {
		t.Fatalf("总数据条数不匹配，期望 %d，实际 %d", expectedCount, len(allUsers))
	}

	fmt.Printf("成功追加数据，总共 %d 条用户数据\n", len(allUsers))
}

// TestProductStruct 测试产品结构体
func TestProductStruct(t *testing.T) {
	processor := NewCSVProcessor("test_products.csv")
	defer processor.DeleteFile()

	// 创建产品数据
	products := []Product{
		{
			ID:          1,
			Name:        "iPhone 15",
			Price:       5999.00,
			Category:    "手机",
			Description: "苹果最新旗舰手机",
		},
		{
			ID:          2,
			Name:        "MacBook Pro",
			Price:       12999.00,
			Category:    "笔记本",
			Description: "专业级笔记本电脑",
		},
		{
			ID:          3,
			Name:        "AirPods Pro",
			Price:       1999.00,
			Category:    "耳机",
			Description: "主动降噪无线耳机",
		},
	}

	// 写入 CSV 文件
	err := processor.WriteFromStruct(products)
	if err != nil {
		t.Fatalf("写入产品 CSV 失败: %v", err)
	}

	// 读取验证
	var readProducts []Product
	err = processor.ReadToStruct(&readProducts)
	if err != nil {
		t.Fatalf("读取产品 CSV 失败: %v", err)
	}

	// 验证数据
	if len(readProducts) != len(products) {
		t.Fatalf("产品数据条数不匹配，期望 %d，实际 %d", len(products), len(readProducts))
	}

	fmt.Printf("成功处理 %d 条产品数据\n", len(readProducts))
}

// TestCSVProcessorMethods 测试 CSV 处理器的其他方法
func TestCSVProcessorMethods(t *testing.T) {
	processor := NewCSVProcessor("test_methods.csv")
	defer processor.DeleteFile()

	// 测试文件路径
	filePath := processor.GetFilePath()
	if filePath != "test_methods.csv" {
		t.Fatalf("文件路径不匹配，期望 test_methods.csv，实际 %s", filePath)
	}

	// 测试文件不存在
	if processor.FileExists() {
		t.Fatal("文件应该不存在")
	}

	// 创建文件
	users := []User{
		{
			ID:    1,
			Name:  "测试用户",
			Age:   25,
			Email: "test@example.com",
			City:  "测试城市",
		},
	}

	err := processor.WriteFromStruct(users)
	if err != nil {
		t.Fatalf("写入测试数据失败: %v", err)
	}

	// 测试文件存在
	if !processor.FileExists() {
		t.Fatal("文件应该存在")
	}

	// 测试文件大小
	size, err := processor.GetFileSize()
	if err != nil {
		t.Fatalf("获取文件大小失败: %v", err)
	}
	if size == 0 {
		t.Fatal("文件大小应该大于 0")
	}

	fmt.Printf("CSV 处理器方法测试通过，文件大小: %d 字节\n", size)
}

// TestMain 测试主函数
func TestMain(m *testing.M) {
	fmt.Println("开始 CSV 处理器测试...")

	// 运行测试
	m.Run()

	fmt.Println("CSV 处理器测试完成")
}
