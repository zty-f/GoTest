package excel

import (
	"bytes"
	"os"
	"testing"
)

// TestExcelReader_FromReader 测试从文件流读取
func TestExcelReader_FromReader(t *testing.T) {
	// 创建测试Excel文件
	testFile := "test_reader.xlsx"
	createTestExcelFile(testFile)
	defer os.Remove(testFile)

	// 读取文件内容到内存
	f, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("打开文件失败: %v", err)
	}
	defer f.Close()

	// 使用文件流读取
	var users []User
	err = NewExcelReaderFromReader(f).
		SetSheet("用户信息").
		ReadToStruct(&users)

	if err != nil {
		t.Fatalf("从文件流读取失败: %v", err)
	}

	if len(users) == 0 {
		t.Error("期望读取到数据，但结果为空")
		return
	}

	// 验证数据
	firstUser := users[0]
	if firstUser.ID != 1 {
		t.Errorf("期望ID=1，实际ID=%d", firstUser.ID)
	}
	if firstUser.Name != "张三" {
		t.Errorf("期望姓名=张三，实际姓名=%s", firstUser.Name)
	}

	t.Logf("从文件流成功读取 %d 条用户数据", len(users))
}

// TestExcelReader_FromBytes 测试从字节流读取（模拟前端上传）
func TestExcelReader_FromBytes(t *testing.T) {
	testFile := "test_bytes.xlsx"
	createTestExcelFile(testFile)
	defer os.Remove(testFile)

	// 读取文件内容为bytes（模拟前端上传的字节流）
	fileBytes, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("读取文件失败: %v", err)
	}

	// 使用bytes.NewReader创建io.ReadSeeker
	reader := bytes.NewReader(fileBytes)

	// 使用文件流读取
	var users []User
	err = NewExcelReaderFromReader(reader).
		SetSheet("用户信息").
		ReadToStruct(&users)

	if err != nil {
		t.Fatalf("从字节流读取失败: %v", err)
	}

	if len(users) == 0 {
		t.Error("期望读取到数据，但结果为空")
		return
	}

	t.Logf("从字节流成功读取 %d 条用户数据", len(users))
	for i, user := range users {
		t.Logf("用户%d: ID=%d, 姓名=%s, 年龄=%d, 薪资=%.2f", i+1, user.ID, user.Name, user.Age, user.Salary)
	}
}

// TestExcelReader_FromBytesComplete 测试文件流的完整功能
func TestExcelReader_FromBytesComplete(t *testing.T) {
	testFile := "test_complete.xlsx"
	createCustomTestExcelFile(testFile)
	defer os.Remove(testFile)

	// 读取为bytes
	fileBytes, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("读取文件失败: %v", err)
	}

	// 从bytes创建reader
	reader := bytes.NewReader(fileBytes)

	// 测试读取产品信息
	var products []Product
	err = NewExcelReaderFromReader(reader).
		SetSheet("产品信息").
		SetHeaderRow(2).    // 表头在第2行
		SetDataStartRow(4). // 数据从第4行开始
		ReadToStruct(&products)

	if err != nil {
		t.Fatalf("从文件流读取完整功能失败: %v", err)
	}

	if len(products) == 0 {
		t.Error("期望读取到产品数据，但结果为空")
		return
	}

	t.Logf("从文件流成功读取 %d 条产品数据", len(products))
	for i, product := range products {
		t.Logf("产品%d: ID=%d, 名称=%s, 价格=%f", i+1, product.ID, product.Name, product.Price)
	}
}

// TestExcelReader_FromReaderMap 测试使用文件流读取为Map
func TestExcelReader_FromReaderMap(t *testing.T) {
	testFile := "test_map.xlsx"
	createTestExcelFile(testFile)
	defer os.Remove(testFile)

	fileBytes, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("读取文件失败: %v", err)
	}

	reader := bytes.NewReader(fileBytes)

	// 使用Map方式读取
	maps, err := NewExcelReaderFromReader(reader).
		SetSheet("用户信息").
		ReadToMap()

	if err != nil {
		t.Fatalf("Map方式读取失败: %v", err)
	}

	if len(maps) == 0 {
		t.Error("期望读取到数据，但结果为空")
		return
	}

	t.Logf("Map方式成功读取 %d 条数据", len(maps))
	for i, m := range maps {
		t.Logf("第%d条: ID=%s, 姓名=%s, 年龄=%s", i+1, m["ID"], m["姓名"], m["年龄"])
	}
}

// TestExcelReader_FromURL 测试从URL读取
func TestExcelReader_FromURL(t *testing.T) {
	// 这个测试需要真实的HTTP服务器
	// 在实际使用中，可以传入任何可访问的Excel文件URL
	t.Log("URL测试需要可访问的HTTP服务器，跳过此测试")

	// 示例用法：
	// var users []User
	// err := NewExcelReaderFromURL("http://example.com/data.xlsx").
	//     SetSheet("用户信息").
	//     ReadToStruct(&users)

	// 如果需要本地测试，可以创建一个简单的HTTP服务器
	// go func() {
	//     http.HandleFunc("/test.xlsx", func(w http.ResponseWriter, r *http.Request) {
	//         http.ServeFile(w, r, "test_file.xlsx")
	//     })
	//     http.ListenAndServe(":8888", nil)
	// }()
	// time.Sleep(time.Second)
	// var users []User
	// err := NewExcelReaderFromURL("http://localhost:8888/test.xlsx").
	//     SetSheet("用户信息").
	//     ReadToStruct(&users)
}

// TestExcelReader_CompareMethods 对比三种读取方式的结果是否一致
func TestExcelReader_CompareMethods(t *testing.T) {
	testFile := "test_compare.xlsx"
	createTestExcelFile(testFile)
	defer os.Remove(testFile)

	// 方法1: 从文件路径读取
	var users1 []User
	err1 := NewExcelReader(testFile).
		SetSheet("用户信息").
		ReadToStruct(&users1)

	// 方法2: 从文件流读取
	fileBytes, _ := os.ReadFile(testFile)
	reader := bytes.NewReader(fileBytes)
	var users2 []User
	err2 := NewExcelReaderFromReader(reader).
		SetSheet("用户信息").
		ReadToStruct(&users2)

	if err1 != nil || err2 != nil {
		t.Fatalf("读取失败: err1=%v, err2=%v", err1, err2)
	}

	// 对比结果
	if len(users1) != len(users2) {
		t.Errorf("数据数量不一致: %d vs %d", len(users1), len(users2))
	}

	for i := range users1 {
		if users1[i].ID != users2[i].ID {
			t.Errorf("用户%d ID不一致: %d vs %d", i, users1[i].ID, users2[i].ID)
		}
		if users1[i].Name != users2[i].Name {
			t.Errorf("用户%d 姓名不一致: %s vs %s", i, users1[i].Name, users2[i].Name)
		}
	}

	t.Logf("两种方法读取结果一致，共 %d 条数据", len(users1))
}
