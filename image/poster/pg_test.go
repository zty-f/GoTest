package poster

import (
	"encoding/base64"
	"fmt"
	"os"
	"testing"
)

// 示例用法
func TestPG(t *testing.T) {
	// // 二维码尺寸 (与前端一致)
	// qrCodeSize := &QRCodeSize{
	// 	Width:  500,
	// 	Height: 500,
	// }
	//
	// // 画布配置 (与前端一致)
	// canvasConfig := &PosterConfig{
	// 	Width:   750,
	// 	Height:  1334,
	// 	QRCodeX: 370,
	// 	QRCodeY: 550,
	// }

	// 创建海报生成器
	generator := NewPosterGenerator(nil, nil)

	// 生成海报示例
	userInfo := &UserInfo{
		ID:                  12345,
		Avatar:              "https://qnyb00.cdn.ipalfish.com/0/img/98/c5/599722f5301deb1d80b957334850",
		Name:                "zruler",
		HasOpenFormalCourse: true,
		IsNovice:            false,
	}

	studyStats := &StudyStatistics{
		DayCount:        23,
		BookCount:       15,
		TotalVocabulary: 365,
	}

	// 生成海报（需要真实的图片 URL）
	base64Str, err := generator.GeneratePosterBase64(
		"https://qnyb00.cdn.ipalfish.com/0/img/22/88/1efb708e68612ac5764f4dbdff23",
		"https://readcamp.cdn.ipalfish.com/readcamp/general/c1/ce/a8a6872c7754a13166cd6e0031cf",
		userInfo,
		studyStats,
		"小明妈妈",
	)
	if err != nil {
		fmt.Printf("Generate poster error: %v\n", err)
		return
	}

	// 根据base64下载图片
	// 根据base64下载图片
	decodedData, err := base64.StdEncoding.DecodeString(base64Str[len("data:image/png;base64,"):])
	if err != nil {
		fmt.Printf("Failed to decode base64 string: %v\n", err)
		return
	}

	file, err := os.Create("poster.png")
	if err != nil {
		fmt.Printf("Failed to create file: %v\n", err)
		return
	}
	defer file.Close()

	if _, err := file.Write(decodedData); err != nil {
		fmt.Printf("Failed to write to file: %v\n", err)
		return
	}

	fmt.Printf("Generated poster base64 (first 100 chars): %s...\n", base64Str[:min(100, len(base64Str))])
}
