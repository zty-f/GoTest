package model

type UserStatus struct {
	// 自增id
	Id int64 `json:"id" form:"id" gorm:"primaryKey" `
	// 用户id
	StuId int64 `json:"stuId" form:"stuId" `
	// 状态类型
	Type string `json:"type" form:"type" `
	// 状态值
	Value string `json:"value" form:"value" `
	// 是否已删除，-1：未删除，1：已删除
	// 创建时间
	CreateTime int64 `json:"createTime" form:"createTime" gorm:"autoCreateTime" `
	// 更新时间
	UpdateTime int64 `json:"updateTime" form:"updateTime" gorm:"autoUpdateTime" `
	// 过期时间
	ExpireTime int64 `json:"expireTime" form:"expireTime" `
}

type LevelInfo struct {
	Level           int             `json:"level"`
	LevelName       string          `json:"levelName"`
	LevelSign       string          `json:"levelSign"`
	LevelValueRange LevelValueRange `json:"levelValueRange"`
}

type LevelValueRange struct {
	Begin int `json:"begin"`
	End   int `json:"end"`
}
