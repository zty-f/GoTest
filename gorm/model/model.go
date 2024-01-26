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
	ExtData    string `gorm:"ext_data" json:"ext_data"`
}

type UserMedal struct {
	Id         int    `gorm:"id" json:"id"`
	StuId      int64  `gorm:"stu_id" json:"stuId"`
	MedalId    int    `gorm:"medal_id" json:"medalId"`
	Year       int    `gorm:"year" json:"year"`
	IsWear     int    `gorm:"is_wear" json:"isWear"`
	ExtData    string `gorm:"ext_data" json:"extData"`
	WearTime   int64  `gorm:"wear_time" json:"wearTime"`
	CreateTime int64  `gorm:"create_time" json:"createTime"`
	NoticeTime int64  `gorm:"notice_time" json:"noticeTime"`
	UpdateTime int64  `gorm:"update_time" json:"updateTime"`
}
