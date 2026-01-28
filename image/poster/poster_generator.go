package poster

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	_ "image/draw"
	"image/png"
	"io"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
)

// PosterConfig 海报配置
type PosterConfig struct {
	Width   int // 海报宽度
	Height  int // 海报高度
	QRCodeX int // 二维码X坐标
	QRCodeY int // 二维码Y坐标
}

// QRCodeSize 二维码尺寸
type QRCodeSize struct {
	Width  int
	Height int
}

// UserInfo 用户信息
type UserInfo struct {
	ID                  int64
	Avatar              string
	Name                string
	HasOpenFormalCourse bool
	IsNovice            bool
}

// StudyStatistics 学习统计
type StudyStatistics struct {
	DayCount        int   // 已读天数
	BookCount       int   // 已读本数
	TotalVocabulary int64 // 累计阅读词汇量
}

// PosterParams 海报生成参数
type PosterParams struct {
	InviterID int64 // 邀请者ID
	PosterID  int64 // 海报ID
	UTMSource int64 // 来源
	TeacherID int64 // 老师ID
	Source    int64 // 来源类型
}

// PosterGenerator 海报生成器
type PosterGenerator struct {
	config     PosterConfig
	qrCodeSize QRCodeSize
	font       font.Face
}

// DefaultPosterConfig 默认海报配置
var DefaultPosterConfig = PosterConfig{
	Width:   750,
	Height:  1334,
	QRCodeX: 370,
	QRCodeY: 550,
}

// DefaultQRCodeSize 默认二维码尺寸
var DefaultQRCodeSize = QRCodeSize{
	Width:  500,
	Height: 500,
}

// NewPosterGenerator 创建海报生成器
func NewPosterGenerator(config *PosterConfig, qrCodeSize *QRCodeSize) *PosterGenerator {
	if config == nil {
		config = &DefaultPosterConfig
	}
	if qrCodeSize == nil {
		qrCodeSize = &DefaultQRCodeSize
	}

	// 加载字体
	fontFace := loadFont(24)

	return &PosterGenerator{
		config:     *config,
		qrCodeSize: *qrCodeSize,
		font:       fontFace,
	}
}

// loadFont 加载字体
func loadFont(size float64) font.Face {
	f, err := truetype.Parse(goregular.TTF)
	if err != nil {
		panic(err)
	}
	face := truetype.NewFace(f, &truetype.Options{
		Size: size,
	})
	return face
}

// GenerateSceneParams 生成场景参数（转换为36进制）
func GenerateSceneParams(params PosterParams) string {
	inviterID := strconv.FormatInt(params.InviterID, 36)
	posterID := strconv.FormatInt(params.PosterID, 36)
	utmSource := strconv.FormatInt(params.UTMSource, 36)
	teacherID := strconv.FormatInt(params.TeacherID, 36)
	source := strconv.FormatInt(params.Source, 36)

	return fmt.Sprintf("%s_%s_%s_%s_%s", inviterID, utmSource, posterID, teacherID, source)
}

// ParseSceneParams 解析场景参数（从36进制转换）
func ParseSceneParams(scene string) (*PosterParams, error) {
	parts := strings.Split(scene, "_")
	if len(parts) != 5 {
		return nil, fmt.Errorf("invalid scene format")
	}

	inviterID, err := strconv.ParseInt(parts[0], 36, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid inviter_id: %v", err)
	}

	utmSource, err := strconv.ParseInt(parts[1], 36, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid utm_source: %v", err)
	}

	posterID, err := strconv.ParseInt(parts[2], 36, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid poster_id: %v", err)
	}

	teacherID, err := strconv.ParseInt(parts[3], 36, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid teacher_id: %v", err)
	}

	source, err := strconv.ParseInt(parts[4], 36, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid source: %v", err)
	}

	return &PosterParams{
		InviterID: inviterID,
		PosterID:  posterID,
		UTMSource: utmSource,
		TeacherID: teacherID,
		Source:    source,
	}, nil
}

// GeneratePoster 生成海报
func (pg *PosterGenerator) GeneratePoster(
	qrCodeURL string,
	backgroundURL string,
	userInfo *UserInfo,
	studyStats *StudyStatistics,
	userName string,
) ([]byte, error) {
	// 创建画布
	dc := gg.NewContext(pg.config.Width, pg.config.Height)

	// 1. 绘制背景图
	bgImg, err := loadImageFromURL(backgroundURL)
	if err != nil {
		return nil, fmt.Errorf("failed to load background image: %v", err)
	}
	dc.DrawImage(bgImg, 0, 0)

	// 2. 绘制用户名（如果有）
	if userName != "" && userInfo != nil {
		if userInfo.HasOpenFormalCourse || !userInfo.IsNovice {
			dc.SetRGB255(85, 85, 85) // #555555
			if err := dc.LoadFontFace("/System/Library/Fonts/STHeiti Light.ttc", 24); err != nil {
				// 如果加载系统字体失败，使用默认字体
				dc.SetFontFace(pg.font)
			}
			dc.DrawString(userName, 450, 386)
		} else if !userInfo.HasOpenFormalCourse {
			dc.SetRGB255(85, 85, 85)
			if err := dc.LoadFontFace("/System/Library/Fonts/STHeiti Light.ttc", 24); err != nil {
				dc.SetFontFace(pg.font)
			}
			dc.DrawString(userName, 450, 402)
		}
	}

	// 3. 绘制用户头像（圆形）
	if userInfo != nil && userInfo.Avatar != "" {
		avatarImg, err := loadImageFromURL(userInfo.Avatar)
		if err == nil {
			var avatarY float64
			var textY float64

			if userInfo.HasOpenFormalCourse || !userInfo.IsNovice {
				avatarY = float64(pg.config.QRCodeY) + 406
				textY = float64(pg.config.QRCodeY) + 426
			} else {
				avatarY = float64(pg.config.QRCodeY) + 480
				textY = float64(pg.config.QRCodeY) + 500
			}

			// 绘制圆形头像
			circleAvatar := makeCircleImage(avatarImg, 38)
			dc.DrawImageAnchored(circleAvatar, 80, int(avatarY), 0.5, 0.5)

			// 绘制用户名
			dc.SetRGBA(1, 1, 1, 1) // 白色
			if err := dc.LoadFontFace("/System/Library/Fonts/STHeiti Light.ttc", 24); err != nil {
				dc.SetFontFace(pg.font)
			}
			dc.DrawString(userInfo.Name, 130, textY)
		}
	}

	// 4. 绘制学习统计数据
	if studyStats != nil && userInfo != nil && (userInfo.HasOpenFormalCourse || !userInfo.IsNovice) {
		// 设置字体颜色
		dc.SetRGB255(51, 51, 51) // #333333

		// 已读天数
		if err := dc.LoadFontFace("/System/Library/Fonts/STHeiti Light.ttc", 26); err != nil {
			dc.SetFontFace(pg.font)
		}
		dc.DrawStringAnchored("已读天数", 140, 1070, 0.5, 0.5)

		if err := dc.LoadFontFace("/System/Library/Fonts/STHeiti Light.ttc", 52); err != nil {
			dc.SetFontFace(loadFont(52))
		}
		dayCountStr := fmt.Sprintf("%d", studyStats.DayCount)
		dc.DrawStringAnchored(dayCountStr, 140, 1130, 0.5, 0.5)

		dayWidth, _ := dc.MeasureString(dayCountStr)
		if err := dc.LoadFontFace("/System/Library/Fonts/STHeiti Light.ttc", 26); err != nil {
			dc.SetFontFace(pg.font)
		}
		dc.DrawString("天", 140+dayWidth/2+10, 1130)

		// 已读本数
		dc.DrawStringAnchored("已读本数", 360, 1070, 0.5, 0.5)

		if err := dc.LoadFontFace("/System/Library/Fonts/STHeiti Light.ttc", 52); err != nil {
			dc.SetFontFace(loadFont(52))
		}
		bookCountStr := fmt.Sprintf("%d", studyStats.BookCount)
		dc.DrawStringAnchored(bookCountStr, 360, 1130, 0.5, 0.5)

		bookWidth, _ := dc.MeasureString(bookCountStr)
		if err := dc.LoadFontFace("/System/Library/Fonts/STHeiti Light.ttc", 26); err != nil {
			dc.SetFontFace(pg.font)
		}
		dc.DrawString("本", 360+bookWidth/2+10, 1130)

		// 累计阅读
		dc.DrawStringAnchored("累计阅读", 580, 1070, 0.5, 0.5)

		if err := dc.LoadFontFace("/System/Library/Fonts/STHeiti Light.ttc", 52); err != nil {
			dc.SetFontFace(loadFont(52))
		}
		totalVocabStr := unitConverter(studyStats.TotalVocabulary)
		dc.DrawStringAnchored(totalVocabStr, 580, 1130, 0.5, 0.5)

		vocabWidth, _ := dc.MeasureString(totalVocabStr)
		if err := dc.LoadFontFace("/System/Library/Fonts/STHeiti Light.ttc", 26); err != nil {
			dc.SetFontFace(pg.font)
		}
		unit := "词"
		if studyStats.TotalVocabulary >= 10000 {
			unit = "万词"
		}
		dc.DrawString(unit, 580+vocabWidth/2+10, 1130)
	}

	// 5. 绘制二维码
	qrImg, err := loadImageFromURL(qrCodeURL)
	if err != nil {
		return nil, fmt.Errorf("failed to load QR code image: %v", err)
	}

	// 缩放二维码到 200x200
	resizedQR := resizeImage(qrImg, 200, 200)

	// 计算二维码位置
	qrX := pg.config.QRCodeX + 175
	qrY := pg.config.QRCodeY + 590

	if pg.qrCodeSize.Width == 0 {
		qrX = pg.config.QRCodeX + 145
		qrY = pg.config.QRCodeY + 490
	}

	dc.DrawImage(resizedQR, qrX, qrY)

	// 导出为 PNG
	var buf bytes.Buffer
	if err := png.Encode(&buf, dc.Image()); err != nil {
		return nil, fmt.Errorf("failed to encode PNG: %v", err)
	}

	return buf.Bytes(), nil
}

// GeneratePosterBase64 生成海报并返回 Base64 字符串
func (pg *PosterGenerator) GeneratePosterBase64(
	qrCodeURL string,
	backgroundURL string,
	userInfo *UserInfo,
	studyStats *StudyStatistics,
	userName string,
) (string, error) {
	data, err := pg.GeneratePoster(qrCodeURL, backgroundURL, userInfo, studyStats, userName)
	if err != nil {
		return "", err
	}

	base64Str := base64.StdEncoding.EncodeToString(data)
	return "data:image/png;base64," + base64Str, nil
}

// loadImageFromURL 从 URL 加载图片
func loadImageFromURL(url string) (image.Image, error) {
	// 如果是 base64 格式
	if strings.HasPrefix(url, "data:image") {
		parts := strings.Split(url, ",")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid base64 image format")
		}
		data, err := base64.StdEncoding.DecodeString(parts[1])
		if err != nil {
			return nil, err
		}
		return png.Decode(bytes.NewReader(data))
	}

	// 从网络加载
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	imgData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(bytes.NewReader(imgData))
	if err != nil {
		return nil, err
	}

	return img, nil
}

// makeCircleImage 将图片裁剪为圆形
func makeCircleImage(src image.Image, radius int) image.Image {
	size := radius * 2
	dc := gg.NewContext(size, size)

	// 绘制圆形裁剪区域
	dc.DrawCircle(float64(radius), float64(radius), float64(radius))
	dc.Clip()

	// 缩放并绘制原图
	srcBounds := src.Bounds()
	srcW := srcBounds.Dx()
	srcH := srcBounds.Dy()

	// 计算缩放比例
	scale := float64(size) / math.Max(float64(srcW), float64(srcH))
	newW := int(float64(srcW) * scale)
	newH := int(float64(srcH) * scale)

	resized := resizeImage(src, newW, newH)
	dc.DrawImageAnchored(resized, radius, radius, 0.5, 0.5)

	return dc.Image()
}

// resizeImage 缩放图片
func resizeImage(src image.Image, width, height int) image.Image {
	dst := image.NewRGBA(image.Rect(0, 0, width, height))

	srcBounds := src.Bounds()
	srcW := srcBounds.Dx()
	srcH := srcBounds.Dy()

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			srcX := x * srcW / width
			srcY := y * srcH / height
			dst.Set(x, y, src.At(srcBounds.Min.X+srcX, srcBounds.Min.Y+srcY))
		}
	}

	return dst
}

// unitConverter 单位转换（处理大数字）
func unitConverter(num int64) string {
	if num == 0 {
		return "0"
	}

	absNum := num
	if num < 0 {
		absNum = -num
	}

	if absNum > 100000000 {
		return fmt.Sprintf("%.2f", float64(num)/100000000)
	} else if absNum > 10000 {
		return fmt.Sprintf("%.2f", float64(num)/10000)
	}

	return fmt.Sprintf("%d", num)
}

// ========== HTTP Handler 示例 ==========

// PosterRequest 海报请求参数
type PosterRequest struct {
	QRCodeURL     string           `json:"qr_code_url"`
	BackgroundURL string           `json:"background_url"`
	UserInfo      *UserInfo        `json:"user_info"`
	StudyStats    *StudyStatistics `json:"study_stats"`
	UserName      string           `json:"user_name"`
	Config        *PosterConfig    `json:"config"`
	QRCodeSize    *QRCodeSize      `json:"qr_code_size"`
}

// PosterResponse 海报响应
type PosterResponse struct {
	Success bool   `json:"success"`
	Data    string `json:"data"` // Base64 编码的图片
	Error   string `json:"error,omitempty"`
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// PosterSceneParams 场景参数（用于二维码）
type PosterSceneParams struct {
	InviterID int64 `json:"inviter_id"` // 邀请者ID
	PosterID  int64 `json:"poster_id"`  // 海报ID
	UTMSource int64 `json:"utm_source"` // 来源标识
	TeacherID int64 `json:"teacher_id"` // 老师ID
	Source    int64 `json:"source"`     // 来源类型
}

// EncodeSceneParams 编码场景参数（转36进制）
func EncodeSceneParams(params PosterSceneParams) string {
	return fmt.Sprintf("%s_%s_%s_%s_%s",
		strconv.FormatInt(params.InviterID, 36),
		strconv.FormatInt(params.UTMSource, 36),
		strconv.FormatInt(params.PosterID, 36),
		strconv.FormatInt(params.TeacherID, 36),
		strconv.FormatInt(params.Source, 36),
	)
}
