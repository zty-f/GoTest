package csv

import (
	"os"
	"testing"
)

// TestData 测试数据结构
type TestData struct {
	OldSalesID string `csv:"原销售id"`
	UID        int64  `csv:"uid"`
	NewSalesID string `csv:"新销售id(如: wurui22393)"`
}

func TestCSVProcessorWithSkip(t *testing.T) {
	// 读取模版.csv文件
	content, err := os.ReadFile("模版.csv")
	if err != nil {
		t.Fatalf("读取文件失败: %v", err)
	}

	t.Logf("文件内容:\n%s", string(content))
	// 测试2: 跳过前3行（标题行+2行备注），只读取数据行
	t.Run("跳过前1行和后2行", func(t *testing.T) {
		processor := NewCSVBytesProcessor2().SetSkipRows(2, 3)
		var data []TestData

		err := processor.ReadFromBytesWithSkip(content, &data)
		if err != nil {
			t.Fatalf("读取CSV失败: %v", err)
		}

		t.Logf("跳过3行后读取到 %d 条数据", len(data))
		for i, item := range data {
			t.Logf("第%d条: OldSalesID=%s, UID=%d, NewSalesID=%s",
				i+1, item.OldSalesID, item.UID, item.NewSalesID)
		}

		// 验证数据
		if len(data) != 2 {
			t.Errorf("期望2条数据，实际得到%d条", len(data))
		}

		if len(data) > 0 {
			// 检查第一条数据
			if data[0].OldSalesID != "1" || data[0].UID != 2 || data[0].NewSalesID != "2" {
				t.Errorf("第一条数据不匹配，期望: OldSalesID=1, UID=2, NewSalesID=2，实际: %+v", data[0])
			}
		}
	})
}
