package poster

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	_ "image/draw"
	"image/png"
	"io"
	_ "math"
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

// DefaultPosterConfig 默认海报配置（与前端defaultSize一致）
var DefaultPosterConfig = PosterConfig{
	Width:   750,
	Height:  1334,
	QRCodeX: 370,
	QRCodeY: 550,
}

// DefaultQRCodeSize 默认二维码尺寸（nil表示不使用特定尺寸）
// 前端默认不传 codeSize 参数，所以这里应该是 0
var DefaultQRCodeSize = QRCodeSize{
	Width:  190,
	Height: 190,
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
	fontFace := loadFont(26)

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

	// 1. 绘制背景图（缩放到画布大小）
	bgImg, err := loadImageFromURL(backgroundURL)
	if err != nil {
		return nil, fmt.Errorf("failed to load background image: %v", err)
	}
	// 前端: cxt.drawImage(bgImg, 0, 0, w, h) - 将背景缩放到画布大小
	resizedBg := resizeImage(bgImg, pg.config.Width, pg.config.Height)
	dc.DrawImage(resizedBg, 0, 0)

	// 2. 绘制用户名（如果有）
	if userName != "" && userInfo != nil {
		dc.SetRGB255(85, 85, 85) // #555555
		if err := dc.LoadFontFace("/System/Library/Fonts/STHeiti Medium.ttc", 24); err != nil {
			// 如果加载系统字体失败，使用默认字体
			dc.SetFontFace(pg.font)
		}

		if userInfo.HasOpenFormalCourse || !userInfo.IsNovice {
			// 前端: fillText(name, 450, 386) - 左对齐，基线坐标
			dc.DrawString(userName, 450, 386)
		} else if !userInfo.HasOpenFormalCourse {
			// 前端: fillText(name, 450, 402) - 左对齐，基线坐标
			dc.DrawString(userName, 450, 402)
		}
	}

	// 3. 绘制用户头像（圆形）
	if userInfo != nil && userInfo.Avatar != "" {
		avatarImg, err := loadImageFromURL(userInfo.Avatar)
		if err == nil {
			var avatarDrawY float64
			var textY float64

			if userInfo.HasOpenFormalCourse || !userInfo.IsNovice {
				// 前端：arc(80, y + 406, 38) 和 drawImage(avatar, 40, y + 370, 80, 80)
				avatarDrawY = float64(pg.config.QRCodeY) + 370 // 绘制起始Y
				textY = float64(pg.config.QRCodeY) + 426
			} else {
				// 前端：arc(80, y + 480, 38) 和 drawImage(avatar, 40, y + 444, 80, 80)
				avatarDrawY = float64(pg.config.QRCodeY) + 444 // 绘制起始Y
				textY = float64(pg.config.QRCodeY) + 500
			}

			// 绘制圆形头像 (80x80)
			circleAvatar := makeCircleImage(avatarImg, 40) // 半径40，直径80
			dc.DrawImage(circleAvatar, 40, int(avatarDrawY))

			// 绘制用户名
			dc.SetRGBA(1, 1, 1, 1) // 白色
			if err := dc.LoadFontFace("/System/Library/Fonts/STHeiti Medium.ttc", 24); err != nil {
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
		if err := dc.LoadFontFace("/System/Library/Fonts/STHeiti Medium.ttc", 26); err != nil {
			dc.SetFontFace(pg.font)
		}
		// 前端: textAlign='center', fillText('已读天数', 140, 1070)
		// 手动计算居中位置：x - width/2
		text1 := "已读天数"
		w1, _ := dc.MeasureString(text1)
		dc.DrawString(text1, 140-w1/2, 1070)

		if err := dc.LoadFontFace("/System/Library/Fonts/STHeiti Medium.ttc", 52); err != nil {
			dc.SetFontFace(loadFont(52))
		}
		dayCountStr := fmt.Sprintf("%d", studyStats.DayCount)
		// 前端: fillText(day_count, 140, 1130) - textAlign仍是center
		dayWidth, _ := dc.MeasureString(dayCountStr)
		dc.DrawString(dayCountStr, 140-dayWidth/2, 1130)

		// 绘制"天"
		if err := dc.LoadFontFace("/System/Library/Fonts/STHeiti Medium.ttc", 26); err != nil {
			dc.SetFontFace(pg.font)
		}
		// 前端: fillText('天', 160 + parseInt(day.width/2), 1130) - textAlign仍是center!
		tianX := 160 + float64(int(dayWidth/2))
		tianW, _ := dc.MeasureString("天")
		dc.DrawString("天", tianX-tianW/2, 1130)

		// 已读本数
		text2 := "已读本数"
		w2, _ := dc.MeasureString(text2)
		dc.DrawString(text2, 360-w2/2, 1070)

		if err := dc.LoadFontFace("/System/Library/Fonts/STHeiti Medium.ttc", 52); err != nil {
			dc.SetFontFace(loadFont(52))
		}
		bookCountStr := fmt.Sprintf("%d", studyStats.BookCount)
		bookWidth, _ := dc.MeasureString(bookCountStr)
		dc.DrawString(bookCountStr, 360-bookWidth/2, 1130)

		if err := dc.LoadFontFace("/System/Library/Fonts/STHeiti Medium.ttc", 26); err != nil {
			dc.SetFontFace(pg.font)
		}
		// 前端: fillText('本', 380 + parseInt(num.width/2), 1130) - textAlign仍是center!
		benX := 380 + float64(int(bookWidth/2))
		benW, _ := dc.MeasureString("本")
		dc.DrawString("本", benX-benW/2, 1130)

		// 累计阅读
		text3 := "累计阅读"
		w3, _ := dc.MeasureString(text3)
		dc.DrawString(text3, 580-w3/2, 1070)

		if err := dc.LoadFontFace("/System/Library/Fonts/STHeiti Medium.ttc", 52); err != nil {
			dc.SetFontFace(loadFont(52))
		}
		totalVocabStr := unitConverter(studyStats.TotalVocabulary)
		vocabWidth, _ := dc.MeasureString(totalVocabStr)
		dc.DrawString(totalVocabStr, 580-vocabWidth/2, 1130)

		if err := dc.LoadFontFace("/System/Library/Fonts/STHeiti Medium.ttc", 26); err != nil {
			dc.SetFontFace(pg.font)
		}
		unit := "词"
		unitX := 600.0 + float64(int(vocabWidth/2))
		if studyStats.TotalVocabulary >= 10000 {
			unit = "万词"
			unitX = 610.0 + float64(int(vocabWidth/2))
		}
		// 前端: fillText('词/万词', 600/610 + parseInt(words.width/2), 1130) - textAlign仍是center!
		unitW, _ := dc.MeasureString(unit)
		dc.DrawString(unit, unitX-unitW/2, 1130)
	}

	// 5. 绘制二维码
	qrImg, err := loadImageFromURL(qrCodeURL)
	if err != nil {
		return nil, fmt.Errorf("failed to load QR code image: %v", err)
	}

	// 缩放二维码到 200x200
	resizedQR := resizeImage(qrImg, pg.qrCodeSize.Width, pg.qrCodeSize.Height)

	// 计算二维码位置
	// 前端: if (codeSize) { x + 175, y + 590 } else { x + 145, y + 490 }
	qrX := pg.config.QRCodeX + 175
	qrY := pg.config.QRCodeY + 590

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

	// 缩放原图到圆形大小
	resized := resizeImage(src, size, size)
	dc.DrawImage(resized, 0, 0)

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
