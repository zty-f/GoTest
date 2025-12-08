package base

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestReadUidsFromTmp(t *testing.T) {
	// 读取 tmp 文件（从项目根目录）
	// Go 测试的工作目录通常是项目根目录（包含 go.mod 的目录）
	tmpPath := "tmp"
	// 如果从 base 目录运行，使用相对路径
	if _, err := os.Stat(tmpPath); os.IsNotExist(err) {
		tmpPath = "../tmp"
	}

	data, err := os.ReadFile(tmpPath)
	if err != nil {
		t.Fatalf("读取文件失败: %v", err)
	}

	// 按行分割
	lines := strings.Split(string(data), "\n")

	// 使用 map 去重
	uidMap := make(map[string]bool)
	var uids []string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		// 如果 uid 不存在，添加到结果中
		if !uidMap[line] {
			uidMap[line] = true
			uids = append(uids, line)
		}
	}

	// 生成逗号分割的字符串
	result := strings.Join(uids, ",")

	fmt.Printf("去重后的 uid 数量: %d\n", len(uids))
	fmt.Printf("逗号分割的字符串: %s\n", result)

	t.Logf("去重后的 uid 数量: %d", len(uids))
	t.Logf("结果字符串长度: %d", len(result))
}
