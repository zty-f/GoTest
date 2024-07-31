package model

import (
	"encoding/json"
	"time"
)

type ImportCheckinQuestionResp struct {
	Filename string                       `json:"filename"`
	Sheets   []ImportCheckinQuestionSheet `json:"sheets"`
}

type ImportCommonItem struct {
	TotalNum   int `json:"totalNum"`
	SuccessNum int `json:"successNum"`
	FailNum    int `json:"failNum"`
}

type ImportCheckinQuestionSheet struct {
	SheetName string `json:"sheetName"`
	ImportCommonItem
}

type GenerateStuCouFlagResp struct {
	TableName string `json:"tableIndex"`
	ImportCommonItem
}

// CheckinQuestion 用户打卡题库表
type CheckinQuestion struct {
	Id          int       `gorm:"column:id" json:"id"`
	Type        int       `gorm:"column:type" json:"type"`
	Content     string    `gorm:"column:content" json:"content"`
	Options     []byte    `gorm:"column:options" json:"options"`
	Answer      string    `gorm:"column:answer" json:"answer"`
	AnswerIndex int       `gorm:"column:answer_index" json:"answer_index"`
	SubjectId   int       `gorm:"column:subject_id" json:"subject_id"`
	Status      int       `gorm:"column:status" json:"status"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type CheckinQuestionOption struct {
	Label   string `json:"label"`
	Content string `json:"content"`
}

func (c *CheckinQuestion) GetOptions() []*CheckinQuestionOption {
	res := make([]*CheckinQuestionOption, 0)
	_ = json.Unmarshal(c.Options, &res)
	return res
}
