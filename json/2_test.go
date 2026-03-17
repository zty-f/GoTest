package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

var (
	StudyTaskIdGetRankExp    = int64(1001) // 获取排行榜经验值任务
	StudyTaskIdReadNewBook   = int64(1002) // 阅读新书籍任务
	StudyTaskIdReadOldBook   = int64(1003) // 阅读旧书籍任务
	StudyTaskIdStudyDuration = int64(1004) // 学习时长任务
)

var taskIdToName = map[int64]string{
	StudyTaskIdGetRankExp:    "获取排行榜经验值",
	StudyTaskIdReadNewBook:   "阅读新书籍",
	StudyTaskIdReadOldBook:   "阅读旧书籍",
	StudyTaskIdStudyDuration: "学习时长",
}

// StudyTaskDetail 学习任务详情
type StudyTaskDetail struct {
	TaskId    int64           `json:"task_id"`             // 任务ID
	Name      string          `json:"name"`                // 任务名称
	Status    int64           `json:"status,omitempty"`    // 任务状态 1-进行中 2-已完成 3-奖励已领取/已发放
	Target    int64           `json:"target"`              // 完成目标次数
	Completed int64           `json:"completed,omitempty"` // 已完成次数
	Reward    *TaskRewardInfo `json:"reward"`              // 任务奖励
}

// TaskRewardInfo 任务奖励信息
type TaskRewardInfo struct {
	RewardType int64  `json:"reward_type"` // 奖励类型 1-星币 2-排行榜经验
	Detail     string `json:"detail"`      // 奖励详情
}

func TestTrans1(t *testing.T) {
	tasks := []*StudyTaskDetail{
		{
			TaskId: StudyTaskIdGetRankExp,
			Name:   taskIdToName[StudyTaskIdGetRankExp],
			Target: 100,
			Reward: &TaskRewardInfo{
				RewardType: 1,
				Detail:     "100",
			},
		},
		{
			TaskId: StudyTaskIdReadNewBook,
			Name:   taskIdToName[StudyTaskIdReadNewBook],
			Target: 5,
			Reward: &TaskRewardInfo{
				RewardType: 2,
				Detail:     "200",
			},
		},
	}
	jsonStr, _ := json.Marshal(tasks)
	fmt.Println(string(jsonStr))
}
