package model

// MedalTypeEnums 奖牌类型枚举
type MedalTypeEnums int

const (
	_                  MedalTypeEnums = iota
	MedalTypeCourse                   // 课程奖牌
	MedalTypeLearnDays                // 学习天数奖牌
	MedalTypeLevel                    // 等级奖牌
	MedalTypeLimited                  // 限定奖牌
)

type MedalTypeItem struct {
	Type       MedalTypeEnums `json:"type"`
	TypeName   string         `json:"typeName"`
	List       []MedalInfo    `json:"list"`
	SeriesList []SeriesMedal  `json:"seriesList"`
}

type SeriesMedal struct {
	SeriesType int         `json:"seriesType"`
	SeriesName string      `json:"seriesName"`
	ChildList  []MedalInfo `json:"childList"`
}

// MedalInfo 奖牌业务相关struct
type MedalInfo struct {
	MedalId        int            `json:"medalId"`
	Name           string         `json:"name"`
	Type           MedalTypeEnums `json:"type"`
	UnlockIcon     string         `json:"unlockIcon"`
	LockIcon       string         `json:"lockIcon"`
	ShareIcon      string         `json:"shareIcon"`
	Model3D        string         `json:"model3D"`
	IsGot          bool           `json:"isGot"`
	GotTime        int64          `json:"gotTime"`
	RuleDesc       string         `json:"ruleDesc"`
	Desc           string         `json:"desc"`
	IsShowProgress bool           `json:"isShowProgress"`
}
