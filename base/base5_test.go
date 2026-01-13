package base

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"golang.org/x/sync/singleflight"
	"log"
	"math"
	"math/rand"
	"regexp"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestNilStruct(t *testing.T) {
	type User struct {
		Name string
	}

	var u *User
	if u == nil {
		t.Log("u is nil")
	} else {
		t.Error("u should be nil")
	}

	u = &User{}

	fmt.Println(u.Name)
	fmt.Println(u.Name)
	fmt.Println(u.Name)
	fmt.Println(cast.ToString(0))

}

func TestThreeMonthAgo(t *testing.T) {
	// 获取当前时间
	now := time.Now()

	// 计算三个月前的时间
	threeMonthsAgo := now.AddDate(0, -3, 0)

	// 打印结果
	fmt.Println("当前时间:", now)
	fmt.Println("三个月前的时间:", threeMonthsAgo)
}

type UserReportConfig struct {
	Reasons []string `json:"reasons"`
}

func TestMarshal(t *testing.T) {
	config := UserReportConfig{
		Reasons: []string{"1", "2", "3"},
	}
	marshal, _ := json.Marshal(config)
	fmt.Printf("%+v\n", string(marshal))
	// {
	//			ReportType: 1,
	//			Reasons:    []string{"垃圾营销", "人身攻击", "淫秽色情", "发布违规内容", "发布其他不适当内容"},
	//		},
	configs := map[int]UserReportConfig{
		1: {
			Reasons: []string{"垃圾营销", "人身攻击", "淫秽色情", "发布违规内容", "发布其他不适当内容"},
		},
	}
	marshal, _ = json.Marshal(configs)
	fmt.Printf("%+v\n", string(marshal))
}

type PrivacyType int64

type PrivacyConfig struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func TestPrivacyTypeMarshal(t *testing.T) {
	config := map[PrivacyType][]*PrivacyConfig{
		1: {
			{Id: 1, Name: "学习时长"},
			{Id: 2, Name: "关注列表"},
			{Id: 3, Name: "关注者列表"},
			{Id: 4, Name: "毕业证书"},
		},
	}
	marshal, _ := json.Marshal(config)
	fmt.Printf("%+v\n", string(marshal))
	x := &PrivacyConfig{}
	fmt.Println(x.Id)
	fmt.Println(x.Name)

	fmt.Println(cast.ToInt64("1234567890")) // 1234567890
	fmt.Println(cast.ToInt64("nh"))         // 0
	fmt.Println(cast.ToInt64("忘记"))         // 0
	dayTime := cast.ToInt64(time.Now().AddDate(0, -1, 0).Format("20060102"))
	fmt.Println(dayTime)
}

func TestMapRange(t *testing.T) {
	m := map[string]int{
		"apple":  1,
		"banana": 2,
		"cherry": 3,
	}
	for key, value := range m {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	for key := range m {
		fmt.Printf("Key: %s\n", key)
	}
	fmt.Println(len(aaa()))
	uidsStr := "3280367737,27293710,3275164075,3280464804,3282575848,3204359228,3275270528,3294291967,22371315,3256764636,3291968063,20769577,3288347355,3131208689,3279487219,3291111357,47892302,2852457406,19342558,3292038465,3289794815,3293372097,3288034035,3288341022,3282349135,3294451576,3251814745,3257901846,2961143459,3294612376,3289047868,3281295627,2894935857,2900411600,3282693355,3294029666,3295221503,3243291770,3294940007,3288833780,3289044209,27049009,3272592297,3241502643,3293370801,44252554,3221595497,3290001068,3282672464,38865913,3292498338,3268675124,3288326699,3257874991,3293394134,3270193239,3295112606,3273001326,3291941143,3288675151,3292359126,3289049279,3280122192,3288791719,3283481633,3282567430,28650955344,3283845641,32793336869,3279082662,3294782914,3245884659,3107403052,3280907674,2934674056,46934422,21200782,3289039041,10738187,3054035056,3282860965,3279321163,2997436510,101728935,3249303996,2930174561,3291967221,3071019513,152875556,28072482,3256544870,90032042,56957812,3273032862,3280657698,3221582966,3288290900,3282656671,3149482286,3272849807,3207674195,3282565970,3279415742,3289250414,3144559074,3282528316,17432552,17411440,3279424093,3289070915,3288830574,3232123025,22027102,3277389113,3289045661,3293786196,26288428,3290358117,3286625793,3282494900,3294136071,97339087,11134388,3289045437,3282567530,23829095,102907758,40191040,3280457754,3282777196,3283773102,3283799211,3292160652,3280993729,30458690,3279767923,3283432728,3293037956,3279695562,3288672852,3260485041,3293307448,3293702144,3289991522,3278606802,3245445148,45031649,3225615203,3274064885,3291944543,3289034563,3293802607,3246625984,2440858416,2926381236,3288885309,7609702,2979760075,3281070077,25763378,2986870656,13874060,2916690436,3293034727,20425291,17307194,70847695,2940279823,328449780,3284726222,3280803562,3282342282,3278830097,54461358,3289045947,2878235825,2889555804,3293736124,3287833640,3294371861,16369991,3254139058,66298630,3283179127,3293441906,2989614214,3282574703,2919921382,3284282878,3281306423,27026574,3290575536,3275165006,3071287117,2994301702,47141174,3282743453,3248661725,3294650631,2892527280,3286899293,3057973994,23905072,3290365134,3286149785,1693143,3293658591,3282393033,3235983377,3050550111,3290459716,3024113897,3281847203,3242910989,3282872196,3242358810,3282952872,3282393033,3024113897,3291668836,3290503630,3293741894,13105918,3280662166,3279369130,3291277281"
	uids := strings.Split(uidsStr, ",")
	fmt.Println(uids)
}

func aaa() []int {
	return nil
}

func TestBase5(t *testing.T) {
	var beginTime, endTime *time.Time
	fmt.Println(beginTime)
	fmt.Println(endTime)
	if beginTime == nil {
		fmt.Println("nil")
	}
	fmt.Println(time.Now().AddDate(0, 0, -1).Unix())
}

func TestGetOaidFromString(t *testing.T) {
	orderContext := `{"oaid":"8888"}} Status:20 Ct:1758855067 Ut:1758855067 Note: EquityOpType:give Childs:[0xc000a17340]`
	oaid := ""
	// 其他广告来源，默认使用oaid,从order_context中解析
	oaidIndex := strings.Index(orderContext, `"oaid"`)
	if oaidIndex != -1 {
		start := oaidIndex + len(`"oaid"`) + 2
		fmt.Println(orderContext[start:])
		end := strings.Index(orderContext[start:], `"`)
		if end != -1 {
			oaid = orderContext[start : start+end]
		}
	}
	fmt.Println(oaid)
	oaid2 := ""
	// 使用正则表达式提取 "oaid" 的值
	re := regexp.MustCompile(`"oaid":"(.*?)"`)
	matches := re.FindStringSubmatch(orderContext)
	if len(matches) > 1 {
		oaid2 = matches[1]
	}

	fmt.Println(oaid2)
}

func TestUnMarshal4(t *testing.T) {
	var config []map[int]PrivacyConfig
	str := `[{
		"1": {"id": 1, "name": "学习时长"},
		"2": {"id": 2, "name": "关注列表"},
		"3": {"id": 3, "name": "关注者列表"},
		"4": {"id": 4, "name": "毕业证书"}
	}]`
	err := json.Unmarshal([]byte(str), &config)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", config)

	// max 函数辅助函数
	maxInt := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	fmt.Println(maxInt(1, 3))
	fmt.Println(maxInt(5, 3))
	fmt.Println(maxInt(6, 3))
	fmt.Println(2 / 1 * 3)
	fmt.Println(2 * 1.00 / 2 * 3)
	fmt.Println(1 * 1.00 / 2 * 3)
	// 四舍五入取整
	fmt.Println(int64(math.Round(float64(60) / float64(31))))
	fmt.Println(int64(math.Round(float64(60) / float64(43))))
	fmt.Println(cast.ToInt64("")) // 0
}

/*
SingleFlight可以将对同一条数据的并发请求进行合并，
只允许一个请求访问数据库中的数据，这个请求获取到的数据结果与其他请求共享。
*/
// SingleFlight
var g singleflight.Group

// getDataFromDB:模拟从数据库中获取key=”my_key”的数据
func getDataFromDB(key string) (string, error) {
	// 使用singleflight.Do ()方法获取数据，仅执行一次
	data, err, _ := g.Do(key, func() (interface{}, error) {
		// 模拟众数据库中获取数据
		log.Printf("get data for key:%s from database", key)
		return "my_data", nil
	})
	if err != nil {
		return "", err
	}
	return data.(string), nil
}
func TestStudent_Speak(t *testing.T) {
	var wg sync.WaitGroup
	reqCount := 1000
	wg.Add(reqCount)
	// 模拟100个并发请求
	for i := 0; i < reqCount; i++ {
		go func() {
			defer wg.Done()
			// 这100个并发请求都希望获取key="my_key"的数据
			data, err := getDataFromDB("my_key")
			if err != nil {
				log.Print(err)
			}
			// 获取数据成功
			log.Printf("I get data:%s for key:my_key", data)
		}()
	}
	wg.Wait()
}

type Conf struct {
	ID  int32 `json:"id"`
	Cnt int64 `json:"cnt"`
}

var conf1s = []Conf{
	{ID: 1, Cnt: 100},
	{ID: 2, Cnt: 200},
	{ID: 3, Cnt: 300},
}

var conf2s = []*Conf{
	{ID: 1, Cnt: 100},
	{ID: 2, Cnt: 200},
	{ID: 3, Cnt: 300},
}

func TestIntStr(t *testing.T) {
	x := int32(3)
	fmt.Println(string(x))        //  Unicode字符
	fmt.Println(cast.ToString(x)) // 3

	a := conf1s[0]
	fmt.Println(a)
	a.Cnt = 99
	fmt.Println(a)
	fmt.Println(conf1s)

	fmt.Println("----------")

	b := conf2s[0]
	fmt.Println(b)
	b.Cnt = 999
	fmt.Println(b)
	fmt.Printf("%+v\n", conf2s[0])

	fmt.Println("----------")

	a1 := conf1s[1]
	a2 := conf1s[1]
	fmt.Printf("%p %p %p\n", &a1.Cnt, &a2.Cnt, &conf1s[1].Cnt) // 地址不相同
	conf1s[1].Cnt = 888
	fmt.Printf("%d %d %d\n", a1.Cnt, a2.Cnt, conf1s[1].Cnt) // 200 200 888 修改互不影响
	b1 := conf2s[1]
	b2 := conf2s[1]
	fmt.Printf("%p %p %p\n", &b1.Cnt, &b2.Cnt, &conf2s[1].Cnt) // 地址相同
	conf2s[1].Cnt = 888
	fmt.Printf("%d %d %d\n", b1.Cnt, b2.Cnt, conf2s[1].Cnt) // 888 888 888 修改相互影响
}

func TestRandom3(t *testing.T) {
	// 随机获取10-120之间的随机数字  rand.Intn(x) 会返回一个[0,x)之间的随机整数
	fmt.Println(rand.Intn(120-10+1) + 10)
	fmt.Println(rand.Intn(120-10+1) + 10)
	fmt.Println(rand.Intn(120-10+1) + 10)
	fmt.Println(rand.Intn(120-10+1) + 10)
	fmt.Println(rand.Intn(120-10+1) + 10)
	fmt.Println(rand.Intn(120-10+1) + 10)
	fmt.Println(rand.Intn(120-10+1) + 10)
	fmt.Println(rand.Intn(120-10+1) + 10)
}

func TestError(t *testing.T) {
	err := errors.Errorf("The callBack is mismatch.")
	fmt.Println(err)
	fmt.Println(err.Error())
	if strings.Contains(err.Error(), "The callBack is mismatch") {
		fmt.Println("contains")
	}
	err2 := errors.New("The callBack is mismatch.")
	fmt.Println(err2)
	fmt.Println(err2.Error())
	if strings.Contains(err2.Error(), "The callBack is mismatch") {
		fmt.Println("contains")
	}
	fmt.Println(cast.ToInt64("1462281144979947520"))

	fmt.Println(54 - 54%10)
	personalWinDayCnt := 54
	fmt.Println(personalWinDayCnt + 10 - personalWinDayCnt%10)
	personalWinDayCnt = 80
	fmt.Println(personalWinDayCnt + 10 - personalWinDayCnt%10)
	personalWinDayCnt = 119
	fmt.Println(personalWinDayCnt + 10 - personalWinDayCnt%10)

	fmt.Println(20250923 / 100)
	fmt.Println(20250123 / 100)
	fmt.Println(20250407 / 100)
	fmt.Println(0 / 100)

	tt, _ := time.ParseInLocation("20060102", "20250923", time.Local)
	fmt.Println(tt.Weekday())
	fmt.Println(int64(tt.Weekday()))
	tt, _ = time.ParseInLocation("20060102", "20251223", time.Local)
	fmt.Println(tt.Weekday())
	fmt.Println(int64(tt.Weekday()))
	fmt.Println(cast.ToInt64(time.Time{}.Format("20060102")))
}

func RandomStr(list []string) string {
	if len(list) == 0 {
		return ""
	}
	return list[rand.Intn(len(list))]
}

func TestRandomStr(t *testing.T) {
	fmt.Println(RandomStr([]string{"a", "b", "c"}))
	fmt.Println(RandomStr([]string{"a", "b", "c"}))
	fmt.Println(RandomStr([]string{"a", "b", "c"}))
	fmt.Println(RandomStr([]string{"a", "b", "c"}))
	fmt.Println(RandomStr([]string{"a", "b", "c"}))
	fmt.Println(RandomStr([]string{"a", "b", "c"}))
	fmt.Println(RandomStr([]string{"a", "b", "c"}))

	var x []*Conf
	x = nil
	for i, conf := range x {
		fmt.Println(i, conf)
	}
}

type ActivityWinConf struct {
	StartTs       int64 `json:"start_ts"`        // 活动开始时间
	MonthRedoLife int64 `json:"month_redo_life"` // 每月可复活次数
	// 连胜目标相关信息 大于最后一个则按照规律计算
	GoalInfos []*GoalInfo `json:"goal_infos"` // 连胜目标相关信息
}

type GoalInfo struct {
	DayCnt    int64 `json:"day_cnt"`    // 连胜目标天数
	AwardType int64 `json:"award_type"` // 奖励类型 1-星币
	AwardCnt  int64 `json:"award_cnt"`  // 奖励数量
}

func TestStrConf(t *testing.T) {
	conf := &ActivityWinConf{
		StartTs:       time.Now().Unix(),
		MonthRedoLife: 2,
		GoalInfos: []*GoalInfo{
			{DayCnt: 7, AwardType: 1, AwardCnt: 10},
			{DayCnt: 14, AwardType: 1, AwardCnt: 15},
			{DayCnt: 30, AwardType: 1, AwardCnt: 30},
		},
	}
	fmt.Printf("%+v\n", conf)
	str, _ := json.Marshal(conf)
	fmt.Println(string(str))
	_ = "{\"start_ts\":1766576993,\"month_redo_life\":2,\"goal_infos\":[{\"day_cnt\":7,\"award_type\":1,\"award_cnt\":10},{\"day_cnt\":14,\"award_type\":1,\"award_cnt\":15},{\"day_cnt\":30,\"award_type\":1,\"award_cnt\":30}]}"

	var newConf *ActivityWinConf
	fmt.Printf("%+v\n", newConf)
}

func TestSplit2(t *testing.T) {
	SnapshotUri := "readcamp/general/2a/d4/45bfa1f3b4ac10f20a50dd1b8fcb"
	uriParts := strings.Split(SnapshotUri, "/")
	curImgMd5 := uriParts[len(uriParts)-1]
	fmt.Println(curImgMd5)

	SnapshotUri2 := "0/img/2a/d4/45bfa1f3b4ac10f20a50dd1b8fcb"
	uriParts2 := strings.Split(SnapshotUri2, "/")
	curImgMd52 := uriParts[len(uriParts2)-1]
	fmt.Println(curImgMd52)

	fmt.Println(curImgMd5 == curImgMd52)
}

type Conf1 struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func TestUnmarshal(t *testing.T) {
	conf := &Conf1{}
	err := json.Unmarshal([]byte(`{"id":123,"name":"zty"}`), conf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", conf)
	SpecialPackageUtmSourceMap := map[int64]string{
		660411800375556: "1000202",
		691049570629888: "1000268",
		691049548658944: "1000267",
		700379330763010: "1000311",
		702027600550148: "1000316",
		714039100485888: "120052",
		717005115625730: "1000350",
		717004896217346: "1000349",
		723085443864832: "1000363",
	}
	b, _ := json.Marshal(SpecialPackageUtmSourceMap)
	fmt.Println(string(b))
}

type ReplyInfo struct {
	Handoff         bool        `json:"handoff,omitempty"` // 是否需要人工介入
	HandoffType     string      `json:"handoff_type,omitempty"`
	Priority        string      `json:"priority,omitempty"`
	RiskLevel       string      `json:"risk_level,omitempty"`
	Tags            []string    `json:"tags,omitempty"`
	Summary         string      `json:"summary,omitempty"` // ai总结用户的提问
	NextAction      string      `json:"next_action,omitempty"`
	SuggestedReply  string      `json:"suggested_reply,omitempty"` // 回复建议
	AgentNotes      string      `json:"agent_notes,omitempty"`
	UserText        string      `json:"user_text,omitempty"`
	StudentId       string      `json:"student_id,omitempty"`
	OrderId         string      `json:"order_id,omitempty"`
	OrderStatus     string      `json:"order_status,omitempty"`
	LogisticsStatus string      `json:"logistics_status,omitempty"`
	DaysAfterPay    interface{} `json:"days_after_pay,omitempty"`
	Channel         string      `json:"channel,omitempty"`
}

func TestMarshalReplyInfo(t *testing.T) {
	replyInfo := &ReplyInfo{
		Handoff:        true,
		Summary:        "This is a summary.",
		NextAction:     "Follow up",
		SuggestedReply: "Suggested reply text.",
	}
	data, err := json.Marshal(replyInfo)
	if err != nil {
		fmt.Println("Error marshaling ReplyInfo:", err)
		return
	}
	fmt.Println("Marshaled ReplyInfo:", string(data))
}
