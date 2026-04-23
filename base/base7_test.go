package base

import (
	"encoding/json"
	"fmt"
	"math"
	"net/url"
	"strings"
	"test/util"
	"testing"
	"time"

	"github.com/spf13/cast"
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

type Data struct {
	A int64 `json:"a"`
	*BigDataDemo
	Ct int64 `json:"ct"`
}

func TestBigDataDemo(t *testing.T) {
	x := Data{
		A: 1,
		BigDataDemo: &BigDataDemo{
			TeacherName:       "teacher1",
			AttendCourseNum:   100,
			AttendCourseRatio: 0.5,
		},
		Ct: Apple,
	}
	marshal, _ := json.Marshal(x)
	fmt.Println(string(marshal))
	y := Data{}
	json.Unmarshal(marshal, &y)
	fmt.Println(y)

	a := int64(3)
	b := int64(4)
	z := a / b
	fmt.Println(z)
	m := float32(a) / float32(b)
	fmt.Println(m)
}

func TestRound(t *testing.T) {
	baseTarget := int64(500)
	difficultyCoef := float64(0.432)
	res := int64(math.Round(float64(baseTarget)*difficultyCoef)) % 10
	fmt.Println(res)
	res2 := cast.ToString(math.Round(float64(baseTarget) * difficultyCoef))
	fmt.Println(res2)
}

type PP struct {
	Age  int64  `json:"age"`
	Name string `json:"name"`
}

func TestPoints(t *testing.T) {
	x := []*PP{
		{
			Age:  18,
			Name: "zty",
		},
		{
			Age:  20,
			Name: "zty2",
		},
	}
	var y *PP
	for _, v := range x {
		if v.Age == 18 {
			y = v
			break
		}
	}
	fmt.Printf("y:%v\n", y)
	y.Name = "zty3"
	fmt.Printf("y:%v\n", y)
	for _, v := range x {
		fmt.Printf("x:%v\n", v)
	}

	split := strings.Split("", ",")
	fmt.Printf("%v\n", split)
}

type Time struct {
	ID         int64     `json:"id"`
	CreateTime time.Time `json:"create_time"`
}

func TestTimeAdd(t *testing.T) {
	today := time.Now().AddDate(0, 0, -2)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(today)
	fmt.Println(time.Time{}.IsZero())
	tt := Time{
		ID: 1,
	}
	fmt.Println(tt.CreateTime.IsZero())

	x := []int{1, 3, 4, 4, 5}
	fmt.Println(x[:5])
	fmt.Println(util.GetDayStartTime(time.Now().AddDate(0, 0, -90)))
}

func FormatResultNum(num int64) int64 {
	if num <= 10 {
		return num
	}
	return num - num%10
}

func TestStudent_Speak2(t *testing.T) {
	fmt.Println(FormatResultNum(1))
	fmt.Println(FormatResultNum(4))
	fmt.Println(FormatResultNum(7))
	fmt.Println(FormatResultNum(123))
	fmt.Println(FormatResultNum(14))
	fmt.Println(FormatResultNum(26))
}

type Mm struct {
	Marks []string `json:"marks"`
}

func TestMm(t *testing.T) {
	x := Mm{
		Marks: []string{"A", "B", "C"},
	}
	marshal, _ := json.Marshal(x)
	fmt.Println(string(marshal))
	fmt.Println(x.Marks)
	ss := "{\"marks\":[\"low_value_users\"]}"
	var y Mm
	err := json.Unmarshal([]byte(ss), &y)
	fmt.Println(err)
	fmt.Println(y.Marks)
	t1 := time.Now()
	t2 := util.GetDayStartTime(time.Now().Add(-30 * 24 * time.Hour))
	fmt.Println(t2)
	fmt.Println(t1)

	query, err := url.ParseQuery("adid=0&oaid=00000000-0000-0000-0000-000000000000&caid=4ff0e6816d3af8e53878b91c720a7cfb_20250325,dbc755a07175c7cd1fc81ea3e15744f5_20230330&timestamp=797412021675&actionType=ACTIVE&ip=0.0.0.0&ipalfish_device_id=879AC2B7-B8D1-4C03-A19C-45A21EF3366D&h_dt=1&channel=appstore")
	fmt.Println(query)
}

func TestTimeAdd2(t *testing.T) {
	x := []int{1, 3, 4, 4, 5}
	fmt.Println(x[:5])
}

func TestSliceReplace(t *testing.T) {
	x := [][]string{
		{""},
		{"1", ""},
		{"2"},
	}
	dempMap := make(map[string]bool)
	for _, v := range x {
		for _, e := range v {
			if _, ok := dempMap[e]; ok {
				fmt.Printf("repeat:%v\n", e)
				continue
			}
			dempMap[e] = true
		}
	}
	fmt.Println(x)
	y := 21
	s := "jdnnfjfj"
	fmt.Println(fmt.Sprintf(s, y))
}

func setZeroes(matrix [][]int) {
	seti := make(map[int]struct{})
	setj := make(map[int]struct{})
	m, n := len(matrix), len(matrix[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				seti[i] = struct{}{}
				setj[j] = struct{}{}
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if _, ok := seti[i]; ok {
				matrix[i][j] = 0
				continue
			}
			if _, ok := setj[j]; ok {
				matrix[i][j] = 0
				continue
			}
		}
	}
}

// s1 = "ttlewaterbottlewaterbo", s2 = "tlewaterbot"
func isFlipedString(s1 string, s2 string) bool {
	return len(s1) == len(s2) && strings.Contains(s1+s1, s2)
}

type ListNode2 struct {
	Val  int
	Next *ListNode2
}

// 输入：[1, 1, 1, 1, 2]
// 输出：[1, 2]
// 输入：[1, 2, 3, 3, 2, 1]
// 输出：[1, 2, 3]  移除重复元素，保留第一个出现的节点
// O(n) 时间复杂度 O(n) 空间复杂度
func removeDuplicateNodes1(head *ListNode2) *ListNode2 {
	set := make(map[int]struct{})
	pre, cur := head, head
	for cur != nil {
		if _, ok := set[cur.Val]; !ok {
			set[cur.Val] = struct{}{}
			pre = cur
		} else {
			pre.Next = cur.Next // 删除cur节点
		}
		cur = cur.Next
	}
	return head
}

// O(n^2) 时间复杂度 O(1) 空间复杂度 两次循环遍历，判断下一个节点是否和当前相同，相同则删除下一个节点
func removeDuplicateNodes2(head *ListNode2) *ListNode2 {
	pre := head
	for pre != nil {
		cur := pre
		for cur.Next != nil {
			if cur.Next.Val == pre.Val {
				cur.Next = cur.Next.Next // 删除cur节点
			} else {
				cur = cur.Next
			}
		}
		pre = pre.Next
	}
	return head
}

// 找出单向链表中倒数第 k 个节点。返回该节点的值。
// 输入： 1->2->3->4->5 和 k = 2
// 输出： 4
// 1、先遍历链表，统计链表长度h
// 2、计算倒数第k个节点的索引x = h - k
// 3、遍历链表，直到索引为x，返回该节点的值
func kthToLast1(head *ListNode2, k int) int {
	h := 0
	cur := head
	for cur != nil {
		h++
		cur = cur.Next
	}
	x := h - k
	for x > 0 {
		head = head.Next
		x--
	}
	return head.Val
}

// 2、快慢指针 快指针先走k步，慢指针再走，快指针走到头时，慢指针就走到倒数第k个节点
func kthToLast2(head *ListNode2, k int) int {
	i, j := head, head
	for k > 0 {
		k--
		i = i.Next
	}
	for i != nil {
		i = i.Next
		j = j.Next
	}
	return j.Val
}
