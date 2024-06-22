package model

type MonthlyReportResponse struct {
	// 基础信息
	StuId       int64  `json:"stuId"`       // 学生Id
	GrowthValue int    `json:"growthValue"` // 成长值
	Level       int    `json:"level"`       // 等级
	LevelName   string `json:"levelName"`   // 等级名称
	Star        int    `json:"star"`        // 等级星星数量

	// 头部信息
	Year      string `json:"year"`      // 年份
	Month     string `json:"month"`     // 月份
	BeginDate string `json:"beginDate"` // 月报统计开始时间
	EndDate   string `json:"endDate"`   // 月报统计结束时间
	Keyword   string `json:"keyword"`   // 本月关键词

	// 战力值
	PowerValue     int              `json:"powerValue"`     // 战力值
	TaskFinishList []TaskFinishInfo `json:"taskFinishList"` // 任务完成次数列表

	// 等级成长信息
	MonthGrowthValue       int                     `json:"monthGrowthValue"`       // 本月新获得的成长值（不含回退）
	GrowthValueOverRate    int                     `json:"growthValueOverRate"`    // 超过人数的百分比
	LevelDesc              string                  `json:"levelDesc"`              // 等级文案，展示升级到或者保持在
	NextLevel              int                     `json:"nextLevel"`              // 下一个等级
	ToNextLevelGrowthValue int                     `json:"toNextLevelGrowthValue"` // 升级到下一级所需的成长值
	LevelPercent           int                     `json:"levelPercent"`           // 升级进度百分比
	ValueChangeList        []MonthStageGrowthValue `json:"valueChangeList"`        // 本月成长值阶段变化列表

	// 金币相关信息
	GoldIncomeNum       int               `json:"goldIncomeNum"`       // 本月用户累计获得的金币数量
	ClassGoldIncomeNum  int               `json:"classGoldIncomeNum"`  // 本月课堂上累计获得金币
	ClassGoldIncomeRate int               `json:"classGoldIncomeRate"` // 课堂金币占比
	LevelGoldIncomeNum  int               `json:"levelGoldIncomeNum"`  // 等级金币
	LevelGoldIncomeRate int               `json:"levelGoldIncomeRate"` // 等级金币占比
	OtherGoldIncomeNum  int               `json:"otherGoldIncomeNum"`  // 其他金币
	OtherGoldIncomeRate int               `json:"otherGoldIncomeRate"` // 其他金币占比
	GoldExpendNum       int               `json:"goldExpendNum"`       // 本月累计使用的金币数
	ExchangeSkuSum      int               `json:"exchangeSkuSum"`      // 统计本月累计兑换的商品数
	ExchangeSkuList     []ExchangeSkuInfo `json:"exchangeSkuList"`     // 金币兑换的商品列表
}

type TaskFinishInfo struct {
	TaskId        int    `json:"taskId"`        // 任务Id
	TaskName      string `json:"taskName"`      // 任务名称
	TaskFinishNum int    `json:"taskFinishNum"` // 任务完成数量
}

type MonthStageGrowthValue struct {
	DateRange string `json:"dateRange"` // 日期范围
	AddValue  int    `json:"addValue"`  // 获得成长值
}

type ExchangeSkuInfo struct {
	SkuId      string `json:"skuId"`      // 商品Id、订单Id
	SkuName    string `json:"skuName"`    // 商品名称、金币抵现
	SkuIcon    string `json:"skuIcon"`    // 商品图链接
	PayPoint   int    `json:"payPoint"`   // 兑换金币数量
	CreateTime int    `json:"createTime"` // 这笔兑换产生的时间（用于排序）
}

// MonthlyReportStatusResponse 月报状态接口响应
type MonthlyReportStatusResponse struct {
	Status int    `json:"status"`
	Url    string `json:"url"`
}
