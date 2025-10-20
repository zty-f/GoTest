package excel

import (
	"os"
	"testing"
	"time"
)

// TestStructExcelWriter_Basic 测试基本的结构体导出
func TestStructExcelWriter_Basic(t *testing.T) {
	// 准备数据
	data := []*User{
		{ID: 1, Name: "张三", Age: 25, Email: "a@a.com", Salary: 5000.5, IsActive: true, Birthday: time.Date(1998, 5, 15, 0, 0, 0, 0, time.Local)},
		{ID: 2, Name: "李四", Age: 30, Email: "b@b.com", Salary: 6000, IsActive: false, Birthday: time.Date(1993, 8, 20, 0, 0, 0, 0, time.Local)},
	}

	file := "export_users.xlsx"
	defer os.Remove(file)

	err := NewStructExcelWriter(data).
		UseTagHeaders().
		TimeFormat("2006-01-02").
		SavePath(file).
		Sheet("用户信息").
		ToExcel().
		Error
	if err != nil {
		t.Fatalf("导出失败: %v", err)
	}
}

// TestStructExcelWriter_CustomColumnsAndHeaders 测试自定义列和标题导出
func TestStructExcelWriter_CustomColumnsAndHeaders(t *testing.T) {
	data := []Product{
		{ID: 1, Name: "A", Category: "C1", Price: 10.5, Stock: 5, Description: "d1", IsAvailable: true},
		{ID: 2, Name: "B", Category: "C2", Price: 20, Stock: 10, Description: "d2", IsAvailable: false},
	}
	file := "export_products.xlsx"
	defer os.Remove(file)

	err := NewStructExcelWriter(data).
		Columns("ID", "Name", "Price", "Stock").
		Headers("产品ID", "产品名称", "价格", "库存").
		SavePath(file).
		Sheet("产品信息").
		ToExcel().
		Error
	if err != nil {
		t.Fatalf("导出失败: %v", err)
	}
}

// TestStructExcelWriter_PointerSlice 测试指针切片导出
func TestStructExcelWriter_PointerSlice(t *testing.T) {
	data := []*Employee{
		{ID: 1, Name: "张经理", Department: "技术部", Position: "经理", Salary: 10000, HireDate: time.Now(), IsManager: true, PhoneNumber: "10086"},
		{ID: 2, Name: "李工程师", Department: "技术部", Position: "高级", Salary: 8000, HireDate: time.Now(), IsManager: false},
	}
	file := "export_employees.xlsx"
	defer os.Remove(file)

	err := NewStructExcelWriter(data).
		UseTagHeaders().
		SavePath(file).
		Sheet("员工信息").
		ToExcel().
		Error
	if err != nil {
		t.Fatalf("导出失败: %v", err)
	}
}
