package main

// XiaohongshuLeadsMessage 小红书私信留资消息结构体
type XiaohongshuLeadsMessage struct {
	Data      XiaohongshuLeadsData `json:"data"`      // 私信留资数据
	Timestamp int64                `json:"timestamp"` // 时间戳
	Source    string               `json:"source"`    // 来源，如"小红书"
}

// XiaohongshuLeadsData 私信留资数据
type XiaohongshuLeadsData struct {
	ID             int      `json:"id"`               // 私信留资ID
	Type           int      `json:"type"`             // 类型，0:新增,1:更新内容，2：更新广告归因信息（计划id/单元id/创意id/笔记）
	Time           string   `json:"time"`             // 私信线索创建时间
	RedID          string   `json:"red_id"`           // 小红书id
	NickName       string   `json:"nick_name"`        // 用户昵称
	Area           string   `json:"area"`             // 省份地区
	LeadsTag       string   `json:"leads_tag"`        // 线索标签，跟进中/留客资/高潜成交/已成单/无意向
	PhoneNum       string   `json:"phone_num"`        // 电话号码
	Wechat         string   `json:"wechat"`           // 微信号
	Remark         string   `json:"remark"`           // 备注信息
	Source         string   `json:"source"`           // 标注客服，专业号后台/留资组件/账户客服名称/-
	AdAccount      string   `json:"ad_account"`       // 投放账户名
	InfoStatus     string   `json:"info_status"`      // 是否留资，未留资/已留资
	AutoRecognize  string   `json:"auto_recognize"`   // 是否自动识别，否/是
	NoteLink       string   `json:"note_link"`        // 笔记链接
	CampaignID     int      `json:"campaign_id"`      // 广告计划ID，广告未归因完成时，默认值为-1
	CampaignName   string   `json:"campaign_name"`    // 广告计划名称，广告未归因完成时，默认值为""
	UnitID         int      `json:"unit_id"`          // 广告单元ID，非广告来源时，默认值为-1
	UnitName       string   `json:"unit_name"`        // 广告单元名称，非广告来源时，默认值为""
	CreativityID   int      `json:"creativity_id"`    // 广告创意ID，非广告来源时，默认值为-1
	CreativityName string   `json:"creativity_name"`  // 创意名称，非广告来源时，默认值为""
	MsgReceiveName string   `json:"msg_receive_name"` // 私信接收人
	MsgReceiveID   string   `json:"msg_receive_id"`   // 私信接收人id
	LinkID         string   `json:"link_id"`          // 获客链接id，默认值为""
	LinkName       string   `json:"link_name"`        // 获客链接名称，默认值为""
	OpenTalk       string   `json:"open_talk"`        // 是否开口，默认值为""
	StaffName      string   `json:"staff_name"`       // 员工姓名，默认值为""
	StaffLabels    []string `json:"staff_labels"`     // 私信接收人标签，默认值为[]
	StaffCountry   string   `json:"staff_country"`    // 私信接收人国家，默认值为""
	StaffProvince  string   `json:"staff_province"`   // 私信接收人省，默认值为""
	StaffCity      string   `json:"staff_city"`       // 私信接收人城市，默认值为""
	ExtInfo        string   `json:"ext_info"`         // 补充信息【服务卡线索】，默认值为""
}
