package model

// UserTask 用户任务记录表
type UserTask struct {
	Id         int    `gorm:"id" json:"id"`
	TaskId     int    `gorm:"task_id" json:"taskId"`
	StuId      int64  `gorm:"stu_id" json:"stuId"`
	CreateTime int    `gorm:"create_time" json:"createTime"`
	UpdateTime int    `gorm:"update_time" json:"updateTime"`
	UniqueId   string `gorm:"unique_id" json:"uniqueId"`
	IsDeleted  int    `gorm:"is_deleted" json:"isDeleted"`
	Extra      string `gorm:"extra" json:"extra"`
}
