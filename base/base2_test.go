package base

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/spf13/cast"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"
)

// 两个时间范围的交集判断
func Test_1(t *testing.T) {
	time1 := time.Now()
	time2 := time.Now().Add(7 * time.Hour * 24)
	beginTime := time.Unix(1726279602, 0)
	endTime := time.Unix(1726379602, 0)
	// 如果time1到time2的时间范围内存在一个时间点在beginTime和endTime之间，就返回true，否则返回false
	// 上述情况归根结底就是两个范围判断交集的问题
	if hasOverlap(time1, time2, beginTime, endTime) {
		fmt.Println("The time ranges overlap")
	} else {
		fmt.Println("The time ranges do not overlap")
	}
}

func hasOverlap(time1, time2, beginTime, endTime time.Time) bool {
	return (time1.Before(endTime) || time1.Equal(endTime)) && (time2.After(beginTime) || time2.Equal(beginTime))
}

func TestCast(t *testing.T) {
	// cast 方法如果不是数值使用cast.ToInt会返回0
	str1 := "20019_finish"
	str2 := "2001000"
	str3 := "finishNum"
	fmt.Println(cast.ToInt(str1)) // 0
	fmt.Println(cast.ToInt(str2)) // 2001000
	fmt.Println(cast.ToInt(str3)) // 0
}

type MapData struct {
	Map  map[string]string
	Name string
	Id   int
}

func TestMapNil(t *testing.T) {
	m := &MapData{}
	m.Map["1"] = "1"
	fmt.Println(m.Map)
}

type Prize struct {
	PrizeId    string `json:"prizeId"`
	PrizeName  string `json:"prizeName"`
	PrizeType  int    `json:"prizeType"`
	PrizeCount int    `json:"prizeCount"`
	PrizeImage string `json:"prizeImage"`
}

func TestModelStr(t *testing.T) {
	prizes := make([]Prize, 0)
	prizes = append(prizes, Prize{
		PrizeId:    "86",
		PrizeName:  "蝙蝠侠",
		PrizeType:  4,
		PrizeCount: 1,
		PrizeImage: "https://static-inc.xiwang.com/mall/caa9c21cb0d27b62365d996753cacb88.png",
	})
	prizes = append(prizes, Prize{
		PrizeId:    "xxx",
		PrizeName:  "抽奖次数",
		PrizeType:  5,
		PrizeCount: 3,
		PrizeImage: "https://static-inc.xiwang.com/mall/b4f35c39ee47240a2db91f579056933d.png",
	})
	marshal, _ := json.Marshal(prizes)
	fmt.Println(string(marshal))
}

func TestSprintf(t *testing.T) {
	println(fmt.Sprintf("user%d init roundNo%d", 1, 2))
	planName := "kskksksks"
	planName2 := "2222【真题易错】222"
	println(strings.Contains(planName, "真题易错"))
	println(strings.Contains(planName2, "真题易错"))

	println(strings.Contains("47,48,67", "48"))
}

type UserInfoModReq struct {
	UserId         string     `json:"user_id"`
	BusinesslineId string     `json:"businessline_id"`
	ModInfo        ModInfo    `json:"mod_info"`
	PersonInfo     PersonInfo `json:"person_info"`
}

type ModInfo struct {
	Status   int    `json:"status"` // 1 正常  2 冻结 3 注销
	PersonId string `json:"person_id"`
}

type PersonInfo struct {
	LastPersonId   string                    `json:"last_person_id"`
	ModPersonId    string                    `json:"mod_person_id"`
	LastPersonInfo map[string]map[string]int `json:"last_person_info"`
	ModPersonInfo  map[string]map[string]int `json:"mod_person_info"`
}

func TestModInfo(t *testing.T) {
	modeInfoStr := `{"user_id":"2100053581","businessline_id":"30","mod_info":{"person_id":"100002222"},"person_info":{"last_person_id":"2100053581","mod_person_id":"100002222","last_person_info":{"2100053581":{"30":2100053581},"100002222":{"40":2100053582}},"mod_person_info":{"2100053581":{},"100002222":{"30":2100053581,"40":2100053582}}}}`

	var userInfoModReq UserInfoModReq
	err := json.Unmarshal([]byte(modeInfoStr), &userInfoModReq)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(userInfoModReq)
	fmt.Println(userInfoModReq.PersonInfo)

	var userInfoModReq1 UserInfoModReq
	err1 := sonic.Unmarshal([]byte(modeInfoStr), &userInfoModReq1)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(userInfoModReq1)
	fmt.Println(userInfoModReq1.PersonInfo)
	for _, v := range userInfoModReq1.PersonInfo.ModPersonInfo {
		fmt.Println(len(v))
	}
	var x map[string]string
	fmt.Println(len(x))
}

func TestContext1(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "key", "value")
	fmt.Println(ctx.Value("key"))

	ctx = context.WithValue(ctx, "key", "value1")
	fmt.Println(ctx.Value("key"))

	ctx1 := context.WithValue(ctx, "key", "value2")
	fmt.Println(ctx1.Value("key"))

	ctx2 := context.WithValue(ctx1, "key1", "value3")
	fmt.Println(ctx2.Value("key"))
	fmt.Println(ctx2.Value("key1"))

}

func TestCurlUserModify(t *testing.T) {
	// 设置请求 URL
	urlStr := "http://userapi.inner.xiwang.com/UserCenter/Users/getUserPersonInfo"

	// 设置请求头
	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	// 设置请求体
	data := url.Values{
		"user_id":         {"2100053558"},
		"filter_status[]": {"2", "3"},
	}

	// 构建请求
	req, err := http.NewRequest("POST", urlStr, bytes.NewBufferString(data.Encode()))
	if err != nil {
		fmt.Println(err)
		return
	}

	// 设置请求头
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}

type ListNode struct {
	Value int
	Next  *ListNode
}

// LeetCode 27 判断回文链表
func TestJudgeList(t *testing.T) {
	list1 := &ListNode{
		Value: 1,
		Next: &ListNode{
			Value: 2,
			Next: &ListNode{
				Value: 3,
				Next: &ListNode{
					Value: 3,
					Next: &ListNode{
						Value: 2,
						Next: &ListNode{
							Value: 1,
						},
					},
				},
			},
		},
	}
	fmt.Println(isCircleList(list1))
	fmt.Println(isPalindrome(list1))
}

func isCircleList(list *ListNode) bool {
	if list == nil || list.Next == nil {
		return true
	}
	tmp, cur := list, list
	var pre *ListNode
	// 1 2 3 4 3 2 1
	for tmp != nil && tmp.Next != nil {
		tmp = tmp.Next.Next
		nxt := cur.Next
		cur.Next = pre
		pre, cur = cur, nxt
	}
	mid := cur
	if tmp != nil {
		mid = cur.Next
	}
	// 3 2 1 4 3 2 1
	for mid != nil && pre != nil {
		if mid.Value != pre.Value {
			return false
		}
		nxt := pre
		mid, pre = mid.Next, pre.Next
		nxt.Next, cur = cur, nxt
	}
	return true
}

func reverseList(head *ListNode) *ListNode {
	// 1 2 3
	var prev, cur *ListNode = nil, head
	for cur != nil {
		nextTmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = nextTmp
	}
	return prev
}

func endOfFirstHalf(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	// 找到前半部分链表的尾节点并反转后半部分链表
	firstHalfEnd := endOfFirstHalf(head)
	secondHalfStart := reverseList(firstHalfEnd.Next)

	// 判断是否回文
	p1 := head
	p2 := secondHalfStart
	result := true
	for result && p2 != nil {
		if p1.Value != p2.Value {
			result = false
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	// 还原链表并返回结果
	firstHalfEnd.Next = reverseList(secondHalfStart)
	return result
}
