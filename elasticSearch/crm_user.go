package es

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"strings"
	"test/util"
	"time"
)

// 电销用户
type ESCRMUser struct {
	UserId          int64 `json:"user_id"`
	TeacherId       int32 `json:"teacher_id"`
	PrivateTime     int64 `json:"private_time"`      // 进入私海的时间戳。仅当 user_type=1 时有效
	AddQwState      int64 `json:"add_qw_state"`      // 加微状态 1：未加微 2:已加微 3:已删微
	TrialStartTs    int64 `json:"trial_start_ts"`    // 上课时间
	Intention       int32 `json:"intention"`         // 用户意向 5:决定购买 4:意向高 3:意向一般 2:暂无意向 1:未接通
	NextFollowUpTs  int64 `json:"next_followup_ts"`  // 下次跟进时间
	FinishTrialTs   int64 `json:"finish_trial_ts"`   // 完课时间,有一节体验课完课
	AttendClassTs   int64 `json:"attend_class_ts"`   // 到课时间
	HasFormalCourse bool  `json:"has_formal_course"` // 有正式课
	HasExtendCourse bool  `json:"has_extend_course"` // 有拓展课
	RefundType      int64 `json:"refund_type"`       // 退费类型 1-未退费 2-部分退 3-全退
	// UploadInvitationScreenshotTs int64                `json:"upload_invitation_screenshot_ts"` // 上传转介绍图片时间
	AddressList       []string `json:"address_list"`         // 地址信息
	UtmSourceIds      []int64  `json:"utm_source_ids"`       // 渠道来源id，不限制level
	Tags              []string `json:"tags"`                 // 标签
	LastBindStudyNum  int64    `json:"last_bind_study_num"`  // 捞取后上课数
	LastBindFinishNum int64    `json:"last_bind_finish_num"` // 捞取后完成数
	TotalStudyNum     int64    `json:"total_study_num"`      // 累计上课数
	TotalFinishNum    int64    `json:"total_finish_num"`     // 累计完课数
	// UserEventInfos               []*ESTelemarketingUserEventInfo `json:"user_event_infos,omitempty"`      // 用户事件

	DropReasonWithTs []string `json:"drop_reason_with_ts"` // 掉海原因-时间
	LatelyDropTs     int64    `json:"lately_drop_ts"`      // 最近掉海时间
	AssignReason     string   `json:"assign_reason"`       // 分配原因

	PicOrderAmt               int64 `json:"pic_order_amt"`                // 绘本支付金额
	RcOrderAmt                int64 `json:"rc_order_amt"`                 // 阅读营支付金额
	AISysCoursePayAmt         int64 `json:"ai_sys_course_pay_amt"`        // ai自拼课
	EngTotalOrderAmt          int64 `json:"eng_total_order_amt"`          // 1v1 支付金额
	TotalPayAmt               int64 `json:"total_pay_amt"`                // 累计站内消费（绘本+阅读营+1v1+ai自拼课）
	BindQwCallNum             int64 `json:"bind_qw_call_num"`             // 绑定后企微外呼次数
	BindQwDuration            int64 `json:"bind_qw_duration"`             // 绑定后企微通话时长：秒
	BindFinishCourseDuration  int64 `json:"bind_finish_course_duration"`  // 绑定后完课时长：秒
	TotalFinishCourseDuration int64 `json:"total_finish_course_duration"` // 累计完课时长：秒
	ReadcampBindCallCount     int64 `json:"readcamp_bind_call_count"`     // 绑定后外呼次数

	ThisMonthPosterNum   int64 `json:"this_month_poster_num"`   // 本月转介绍参加次数
	LastMonthPosterNum   int64 `json:"last_month_poster_num"`   // 上月转介绍参加次数
	ThisWeekPosterStatus int64 `json:"this_week_poster_status"` // 本周转介绍海报状态 0:海报已生成,1:已上传(即待审核),2:审核通过,3:审核不通过;-1:其他
}

type ESCRMUsers []*ESCRMUser

func (e ESCRMUsers) UserIds() []int64 {
	var res []int64
	for _, user := range e {
		res = append(res, user.UserId)
	}
	return res
}

// 获取转化情况 1-未转化 2-已转化 3-部分退 4-全额退
func (e ESCRMUser) GetTransformType() int64 {
	if !e.HasFormalCourse && !e.HasExtendCourse && e.RefundType != 3 {
		return 1
	}
	if e.RefundType == 2 {
		return 3
	}
	if e.RefundType == 3 {
		return 4
	}
	// 已转化
	if (e.HasFormalCourse || e.HasExtendCourse) && e.RefundType == 1 {
		return 2
	}
	return 0
}

// IsTransform 获取转化情况
func (e ESCRMUser) IsTransform() bool {
	if e.HasFormalCourse || e.HasExtendCourse {
		return true
	}
	return false
}

type SearchCRMUserRequest struct {
	UserId              int64
	UserIds             []int64
	TeacherId           int32
	TeacherIds          []int64
	AssignTsRange       []int64 // 分配时间区间
	NextFollowUpTsRange []int64 // 下次跟进时间区间
	AttendClassTsRange  []int64 // 到课时间区间
	TrialStartTsRange   []int64 // 上课时间

	IsPaidUser               int32 // 付费用户  1是付费 2是未付费
	AddQwState               int64 // 加微状态 1：未加微 2:已加微 3:已删微
	WebCallAllFailedMaxCount int32 // 外呼均未接通的情况 暂用做 用户未接通<3

	RefundType    int64 // 退费类型 1-未退费 2-部分退 3-全退
	NotRefundType int64 // 退费类型 1-未退费 2-部分退 3-全退
	// UploadInvitationScreenshotTsLt int64    // 本周转介绍截图上传时间
	Intention         []int64  // 意向度
	AddressList       []string // 地址信息
	QueryTag          string   // 标签
	UtmSourceIds      []int64  // 渠道来源
	LastBindStudyNum  int64    // 捞取后上课数
	LastBindFinishNum int64    // 捞取后完课数
	TotalStudyNum     int64    // 累计上课数
	TotalFinishNum    int64    // 累计完课数
	BindFinishCourse  bool     // 绑定后完课时长

	LatelyDropTsRange []int64 // 最近掉海时间

	ReadcampAccCallCountRange        []int64     // 阅读营累计拨打次数
	ReadcampAccCallDurationSecsRange []int64     // 阅读营累计通话时长(秒)
	UserEventType                    int64       // 用户事件
	UserEventTsRange                 []time.Time // 用户事件时间

	AISysCoursePayAmtRange []int64 // ai自拼课
	PicOrderAmtRange       []int64 // 绘本支付金额
	RcOrderAmtRange        []int64 // 阅读营支付金额
	EngTotalOrderAmtRange  []int64 // 1v1 支付金额
	TotalPayAmtRange       []int64 // 累计站内消费（绘本+阅读营+1v1+ai自拼课）

	UserIdGt           int64
	OrderByField       string // 分配时间 private  意向度 intention
	OrderByIncludeZero bool   // 排序包含0
	Asc                bool
	Limit              int64
	Offset             int64
	AssignReason       string // 分配原因

	// 0 也要用来查询，用其他默认值判断是否添加筛选 对其他查询有影响，只能用指针判断
	ThisMonthPosterNum   *int64
	LastMonthPosterNum   *int64
	ThisWeekPosterStatus *int64
}

// SearchCRMUser
/*
Es查询方式第二种：通过elastic包官方方法构建，然后转为json查询即可
*/
func SearchCRMUser(ctx context.Context, req *SearchCRMUserRequest) (users ESCRMUsers, offset int64, total int64, more bool, err error) {

	var search EsSearch
	search.AddMustQuery(elastic.NewTermQuery("user_type", 1))
	if req.UserId > 0 {
		search.AddMustQuery(elastic.NewTermsQuery("user_id", req.UserId))
	}
	if len(req.UserIds) > 0 {
		uids := util.SliceToAny(req.UserIds)
		search.AddMustQuery(elastic.NewTermsQuery("user_id", uids...))
	}
	if req.UserIdGt != 0 {
		search.AddMustQuery(elastic.NewRangeQuery("user_id").Gt(req.UserIdGt))
	}
	if req.TeacherId > 0 {
		search.AddMustQuery(elastic.NewTermsQuery("teacher_id", req.TeacherId))
	} else if len(req.TeacherIds) > 0 {
		tids := util.SliceToAny(req.TeacherIds)
		search.AddMustQuery(elastic.NewTermsQuery("teacher_id", tids...))
	}
	if req.OrderByField != "" {
		search.Sorters = append(search.Sorters, elastic.NewFieldSort(req.OrderByField).Order(req.Asc))
		if req.OrderByIncludeZero == false {
			search.AddMustQuery(elastic.NewRangeQuery(req.OrderByField).Gt(0))
		}
	}

	if req.LastBindStudyNum > 0 {
		search.AddMustQuery(elastic.NewTermsQuery("last_bind_study_num", req.LastBindStudyNum))
	}
	if req.LastBindFinishNum > 0 {
		search.AddMustQuery(elastic.NewRangeQuery("last_bind_finish_num").Gt(0))
	} else if req.LastBindFinishNum == -1 {
		search.AddMustQuery(elastic.NewTermQuery("last_bind_finish_num", 0))
	}
	if req.TotalStudyNum > 0 {
		search.AddMustQuery(elastic.NewTermsQuery("total_study_num", req.TotalStudyNum))
	}
	if req.TotalFinishNum > 0 {
		search.AddMustQuery(elastic.NewRangeQuery("total_finish_num").Gt(0))
	} else if req.TotalFinishNum == -1 {
		search.AddMustQuery(elastic.NewTermQuery("total_finish_num", 0))
	}
	// 通话信息
	if len(req.ReadcampAccCallCountRange) > 0 {
		if req.ReadcampAccCallCountRange[0] != -1 {
			// search.AddMustQuery(elastic.NewRangeQuery("readcamp_acc_call_count").Gte(req.ReadcampAccCallCountRange[0]))
			search.AddMustQuery(elastic.NewNestedQuery("metrics", elastic.NewRangeQuery("metrics.readcamp_acc_call_count").Gte(req.ReadcampAccCallCountRange[0])))
		}
		if req.ReadcampAccCallCountRange[1] != -1 {
			// search.AddMustQuery(elastic.NewRangeQuery("readcamp_acc_call_count").Lte(req.ReadcampAccCallCountRange[1]))
			search.AddMustQuery(elastic.NewNestedQuery("metrics", elastic.NewRangeQuery("metrics.readcamp_acc_call_count").Lte(req.ReadcampAccCallCountRange[1])))
		}
	}
	if len(req.ReadcampAccCallDurationSecsRange) > 0 {
		if req.ReadcampAccCallDurationSecsRange[0] != -1 {
			// search.AddMustQuery(elastic.NewRangeQuery("readcamp_acc_call_duration_secs").Gte(req.ReadcampAccCallDurationSecsRange[0]))
			search.AddMustQuery(elastic.NewNestedQuery("metrics", elastic.NewRangeQuery("metrics.readcamp_acc_call_duration_secs").Gte(req.ReadcampAccCallDurationSecsRange[0])))
		}
		if req.ReadcampAccCallDurationSecsRange[1] != -1 {
			// search.AddMustQuery(elastic.NewRangeQuery("readcamp_acc_call_duration_secs").Lte(req.ReadcampAccCallDurationSecsRange[1]))
			search.AddMustQuery(elastic.NewNestedQuery("metrics", elastic.NewRangeQuery("metrics.readcamp_acc_call_duration_secs").Lte(req.ReadcampAccCallDurationSecsRange[1])))
		}
	}

	// 消费相关
	if len(req.PicOrderAmtRange) > 0 {
		if req.PicOrderAmtRange[0] != -1 {
			search.AddMustQuery(elastic.NewRangeQuery("pic_order_amt").Gte(req.PicOrderAmtRange[0]))
		}
		if req.PicOrderAmtRange[1] != -1 {
			search.AddMustQuery(elastic.NewRangeQuery("pic_order_amt").Lte(req.PicOrderAmtRange[1]))
		}
	}
	if len(req.AISysCoursePayAmtRange) > 0 {
		if req.AISysCoursePayAmtRange[0] != -1 {
			search.AddMustQuery(elastic.NewRangeQuery("ai_sys_course_pay_amt").Gte(req.AISysCoursePayAmtRange[0]))
		}
		if req.AISysCoursePayAmtRange[1] != -1 {
			search.AddMustQuery(elastic.NewRangeQuery("ai_sys_course_pay_amt").Lte(req.AISysCoursePayAmtRange[1]))
		}
	}
	if len(req.RcOrderAmtRange) > 0 {
		if req.RcOrderAmtRange[0] != -1 {
			search.AddMustQuery(elastic.NewRangeQuery("rc_order_amt").Gte(req.RcOrderAmtRange[0]))
		}
		if req.RcOrderAmtRange[1] != -1 {
			search.AddMustQuery(elastic.NewRangeQuery("rc_order_amt").Lte(req.RcOrderAmtRange[1]))
		}
	}
	if len(req.EngTotalOrderAmtRange) > 0 {
		if req.EngTotalOrderAmtRange[0] != -1 {
			search.AddMustQuery(elastic.NewRangeQuery("eng_total_order_amt").Gte(req.EngTotalOrderAmtRange[0]))
		}
		if req.EngTotalOrderAmtRange[1] != -1 {
			search.AddMustQuery(elastic.NewRangeQuery("eng_total_order_amt").Lte(req.EngTotalOrderAmtRange[1]))
		}
	}
	if len(req.TotalPayAmtRange) > 0 {
		if req.TotalPayAmtRange[0] != -1 {
			search.AddMustQuery(elastic.NewRangeQuery("eng_total_order_amt").Gte(req.TotalPayAmtRange[0]))
		}
		if req.TotalPayAmtRange[1] != -1 {
			search.AddMustQuery(elastic.NewRangeQuery("eng_total_order_amt").Lte(req.TotalPayAmtRange[1]))
		}
	}

	if req.AddQwState != 0 {
		search.AddMustQuery(elastic.NewTermQuery("add_qw_state", req.AddQwState))
	}
	if req.IsPaidUser != 0 {
		if req.IsPaidUser == 1 {
			search.AddMustQuery(elastic.NewBoolQuery().
				Should(elastic.NewTermQuery("has_formal_course", true)).
				Should(elastic.NewTermQuery("has_extend_course", true)))
		} else {
			search.AddMustQuery(elastic.NewBoolQuery().
				Should(
					elastic.NewBoolQuery().
						MustNot(elastic.NewExistsQuery("has_formal_course")).
						MustNot(elastic.NewExistsQuery("has_extend_course")),
				).
				Should(
					elastic.NewBoolQuery().
						Must(elastic.NewTermQuery("has_formal_course", false)).
						Must(elastic.NewTermQuery("has_extend_course", false)),
				).
				MinimumShouldMatch("1"),
			)
		}
	}

	if len(req.AssignTsRange) > 0 {
		tsRange := req.AssignTsRange
		search.AddMustQuery(elastic.NewRangeQuery("private_time").Gte(tsRange[0]).Lte(tsRange[1]))
	}
	if len(req.NextFollowUpTsRange) > 0 {
		tsRange := req.NextFollowUpTsRange
		search.AddMustQuery(elastic.NewRangeQuery("next_followup_ts").Gte(tsRange[0]).Lte(tsRange[1]))
		search.AddMustQuery(elastic.NewRangeQuery("next_followup_ts"))
	}
	if len(req.AttendClassTsRange) > 0 {
		tsRange := req.AttendClassTsRange
		search.AddMustQuery(elastic.NewRangeQuery("attend_class_ts").Gte(tsRange[0]).Lte(tsRange[1]))
		search.AddMustQuery(elastic.NewScriptQuery(elastic.NewScript("doc['attend_class_ts'].value > doc['private_time'].value")))
	}

	if req.BindFinishCourse {
		search.AddMustQuery(elastic.NewRangeQuery("bind_finish_course_duration").Gt(0))
	}
	if len(req.TrialStartTsRange) > 0 {
		tsRange := req.TrialStartTsRange
		search.AddMustQuery(elastic.NewRangeQuery("trial_start_ts").Gte(tsRange[0]).Lte(tsRange[1]))
	}

	if req.RefundType > 0 {
		search.AddMustQuery(elastic.NewTermQuery("refund_type", req.RefundType))
	}
	if req.NotRefundType > 0 {
		search.AddMustNotQuery(elastic.NewTermQuery("refund_type", req.NotRefundType))
	}
	// if req.UploadInvitationScreenshotTsLt > 0 {
	//	search.AddMustQuery(elastic.NewBoolQuery().
	//		Should(
	//			elastic.NewRangeQuery("upload_invitation_screenshot_ts").Lt(req.UploadInvitationScreenshotTsLt),
	//			elastic.NewBoolQuery().MustNot(elastic.NewExistsQuery("upload_invitation_screenshot_ts")),
	//		))
	// }
	if len(req.Intention) > 0 {
		var intentions []interface{}
		for _, intention := range req.Intention {
			intentions = append(intentions, intention)
		}
		search.AddMustQuery(elastic.NewTermsQuery("intention", intentions...))
	}
	if len(req.AddressList) > 0 {
		var addressList []interface{}
		for _, addr := range req.AddressList {
			addressList = append(addressList, addr)
		}
		search.AddMustQuery(elastic.NewTermsQuery("address_list", addressList...))
	}
	if req.QueryTag == "no_tags" {
		search.AddMustNotQuery(elastic.NewExistsQuery("tags"))
	} else if strings.TrimSpace(req.QueryTag) != "" {
		search.AddMustQuery(elastic.NewTermQuery("tags", req.QueryTag))
	}

	if len(req.UtmSourceIds) > 0 {
		var utmIds []interface{}
		for _, utmId := range req.UtmSourceIds {
			utmIds = append(utmIds, utmId)
		}
		search.AddMustQuery(elastic.NewTermsQuery("utm_source_ids", utmIds...))
	}

	if req.WebCallAllFailedMaxCount > 0 {
		search.AddMustQuery(elastic.NewRangeQuery("readcamp_bind_call_count").Lte(req.WebCallAllFailedMaxCount))
	}
	if len(req.LatelyDropTsRange) > 1 {
		search.AddMustQuery(elastic.NewRangeQuery("lately_drop_ts").Gte(req.LatelyDropTsRange[0]).Lte(req.LatelyDropTsRange[1]))
	}

	// 用户事件
	if req.UserEventType != 0 {
		search.AddMustQuery(elastic.NewNestedQuery("user_event_infos", elastic.NewTermQuery("user_event_infos.event_type", req.UserEventType)))
	}
	if len(req.UserEventTsRange) > 0 && !req.UserEventTsRange[0].IsZero() {
		search.AddMustQuery(elastic.NewNestedQuery("user_event_infos", elastic.NewRangeQuery("user_event_infos.event_ts").Gte(req.UserEventTsRange[0].Format(time.DateTime))))
	}
	if len(req.UserEventTsRange) > 1 && !req.UserEventTsRange[1].IsZero() {
		search.AddMustQuery(elastic.NewNestedQuery("user_event_infos", elastic.NewRangeQuery("user_event_infos.event_ts").Lte(req.UserEventTsRange[1].Format(time.DateTime))))
	}
	if req.AssignReason != "" {
		search.AddMustQuery(elastic.NewTermQuery("assign_reason", req.AssignReason))
	}
	// 转介绍相关
	if req.LastMonthPosterNum != nil {
		search.AddMustQuery(elastic.NewTermQuery("last_month_poster_num", *req.LastMonthPosterNum))
	} else if req.ThisMonthPosterNum != nil {
		search.AddMustQuery(elastic.NewTermQuery("this_month_poster_num", *req.ThisMonthPosterNum))
	}
	if req.ThisWeekPosterStatus != nil {
		search.AddMustQuery(elastic.NewTermQuery("this_week_poster_status", *req.ThisWeekPosterStatus))
	}

	dsl, err := search.BoolQueryDSL()
	if err != nil {
		return nil, 0, 0, false, err
	}

	var res ESCRMUsers
	data, err := SearchByDSL(ctx, "BussTypeTelemarkingUser", string(dsl), req.Offset, req.Limit)
	if err != nil {
		return res, 0, 0, false, err
	}

	docs := data.Docs
	err = json.Unmarshal([]byte(docs), &res)
	if err != nil {
		return res, 0, 0, false, err
	}
	return res, data.Offset, data.Total, data.More, nil
}

// SearchAggregationByTid 统计每个teacher_id库容
/*
Es 聚合查询，统计每个teacher_id库容
*/
func SearchAggregationByTid(ctx context.Context, req *SearchCRMUserRequest) (map[int64]int64, error) {
	var search EsSearch
	search.AddMustQuery(elastic.NewTermQuery("user_type", 1))
	if req.TeacherId > 0 {
		search.AddMustQuery(elastic.NewTermsQuery("teacher_id", req.TeacherId))
	}
	if req.IsPaidUser != 0 {
		if req.IsPaidUser == 1 {
			search.AddMustQuery(elastic.NewBoolQuery().
				Should(elastic.NewTermQuery("has_formal_course", true)).
				Should(elastic.NewTermQuery("has_extend_course", true)))
		} else {
			search.AddMustQuery(elastic.NewBoolQuery().
				Should(
					elastic.NewBoolQuery().
						MustNot(elastic.NewExistsQuery("has_formal_course")).
						MustNot(elastic.NewExistsQuery("has_extend_course")),
				).
				Should(
					elastic.NewBoolQuery().
						Must(elastic.NewTermQuery("has_formal_course", false)).
						Must(elastic.NewTermQuery("has_extend_course", false)),
				).
				MinimumShouldMatch("1"),
			)
		}
	}
	if req.NotRefundType > 0 {
		search.AddMustNotQuery(elastic.NewTermQuery("refund_type", req.NotRefundType))
	}

	search.Aggregations = append(search.Aggregations, ESSearchAgg{
		Name: "group_by_date_ts",
		Agg:  elastic.NewTermsAggregation().Field("teacher_id").OrderByKey(false).Size(500),
	})

	size := int(req.Limit)
	search.Size = &size

	dsl, err := search.BoolQueryDSL()
	if err != nil {
		return nil, err
	}

	var esRes EsRespBody
	var res = make(map[int64]int64)

	data, err := SearchQueryDslBasic(ctx, "BussTypeTelemarkingUser", string(dsl))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(data), &esRes)
	if err != nil {
		return nil, err
	}
	if len(esRes.Aggregations.GroupByDateTs.Buckets) == 0 {
		return res, nil
	}

	for _, bucket := range esRes.Aggregations.GroupByDateTs.Buckets {
		res[bucket.Key] = bucket.DocCount
	}
	return res, nil
}
