package base

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestDebug(t *testing.T) {
	userIdToDepartment := map[string]int64{}
	useridMap := make(map[string]bool)
	fmt.Println(useridMap)
	fmt.Println(userIdToDepartment)
	today := time.Now().Unix()
	// 时间往前提一分钟 防止有漏掉的情况
	fromDateTime := today
	toDateTime := 1770825600 + 6*24*3600

	fmt.Println(today)
	fmt.Println(fromDateTime)
	fmt.Println(toDateTime)
}

type BigDataDemo struct {
	TeacherName            string  `json:"teacher_name"`              // 教师账户
	AttendCourseNum        int64   `json:"attend_course_num"`         // 参与课程用户数
	AttendCourseRatio      float64 `json:"attend_course_ratio"`       // 参与课程用户数占比
	FinishCourseNum        int64   `json:"finish_course_num"`         // 完成课程用户数
	FinishCourseRatio      float64 `json:"finish_course_ratio"`       // 完成课程用户数占比
	AddWeChatNum24h        int64   `json:"add_wechat_num_24h"`        // 24小时加微数
	AddWeChatNum24hRatio   float64 `json:"add_wechat_num_24h_ratio"`  // 24小时加微数占比
	SetLevelNum24h         int64   `json:"set_level_num_24h"`         // 24小时定级数
	SetLevelNum24hRatio    float64 `json:"set_level_num_24h_ratio"`   // 24小时定级数占比
	SendBookNum48h         int64   `json:"send_book_num_48h"`         // 48小时发书数
	SendBookNum48hRatio    float64 `json:"send_book_num_48h_ratio"`   // 48小时发书数占比
	PosterPassNum          int64   `json:"poster_pass_num"`           // 海报审核通过数
	PosterPassRatio        float64 `json:"poster_pass_ratio"`         // 海报审核通过数占比
	ReferralGetCourseNum   int64   `json:"referral_get_course_num"`   // 转介绍领课数
	ReferralGetCourseRatio float64 `json:"referral_get_course_ratio"` // 转介绍领课数占比
	FirstRefundNum         int64   `json:"first_refund_num"`          // 首单退款用户数
	FirstRefundNumRatio    float64 `json:"first_refund_num_ratio"`    // 首单退款用户数占比
	Reply10MinRatio        float64 `json:"reply_10min_ratio"`         // 10分钟内回复率
	Reply30MinRatio        float64 `json:"reply_30min_ratio"`         // 30分钟内回复率
}

func TestBigDataDemo(t *testing.T) {
	x := BigDataDemo{}
	marshal, _ := json.Marshal(x)
	fmt.Println(string(marshal))

}
