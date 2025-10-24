package main

// API响应结构体
type APIResponse struct {
	Ret  int `json:"ret"`
	Code int `json:"code"`
	Data struct {
		Ent struct {
			BookInfo struct {
				BookID       int64    `json:"bookid"`
				Title        string   `json:"title"`
				Cover        Cover    `json:"cover"`
				Difficulty   int      `json:"difficulty"`
				Tags         []string `json:"tags"`
				Introduction string   `json:"introduction"`
			} `json:"bookinfo"`
			PageInfos []PageInfo `json:"pageinfos"`
		} `json:"ent"`
	} `json:"data"`
}

// 封面信息
type Cover struct {
	Tiny   string `json:"tiny"`
	Origin string `json:"origin"`
	Width  int    `json:"w"`
	Height int    `json:"h"`
}

// 页面信息
type PageInfo struct {
	PageID                   int64  `json:"pageid"`
	ImageURL                 string `json:"imageurl"`
	ListenText               string `json:"listentext"`
	ListenAudioURL           string `json:"listenaudiourl"`
	ExplainAudioURL          string `json:"explainaudiourl"`
	RecordText               string `json:"recordtext"`
	RecordAudioURL           string `json:"recordaudiourl"`
	Translation              string `json:"translation"`
	HasTeach                 bool   `json:"hasteach"`
	HasRecord                bool   `json:"hasrecord"`
	Index                    int    `json:"index"`
	ListenTextPinyin         string `json:"listentext_pinyin"`
	RecordTextPinyin         string `json:"recordtext_pinyin"`
	IsOpenLP                 bool   `json:"is_open_lp"`
	IsOpenRP                 bool   `json:"is_open_rp"`
	BGPicture                string `json:"bg_picture"`
	ListenOfficialEvaluation string `json:"listenofficialevaluation"`
	RecordOfficialEvaluation string `json:"recordofficialevaluation"`
	ScoreText                string `json:"scoretext"`
	Picture                  struct {
		Tiny   string `json:"tiny"`
		Origin string `json:"origin"`
		Width  int    `json:"w"`
		Height int    `json:"h"`
	} `json:"picture"`
}

// Excel导出用的页面数据
type ExcelPageData struct {
	Index          int    `excel:"页码"`
	BGPicture      string `excel:"绘本内容"`
	ListenText     string `excel:"英文字幕原文"`
	Translation    string `excel:"中文翻译原文"`
	ListenAudioURL string `excel:"英文字母音频"`
}

// 配置结构体
type Config struct {
	BookIDs   []int64 `json:"book_ids"`
	OutputDir string  `json:"output_dir"`
	ExcelFile string  `json:"excel_file"`
	APIURL    string  `json:"api_url"`
	Cookie    string  `json:"cookie"`
}

// 下载任务
type DownloadTask struct {
	URL      string
	FilePath string
	Type     string // "image" or "audio"
}
