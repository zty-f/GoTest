package util

import (
	"github.com/spf13/cast"
	"strings"
	"time"
)

var (
	TimeLayoutMap = map[string]string{
		"y": "2006",
		"m": "2006-01",
		"d": "2006-01-02",
		"h": "2006-01-02 15",
		"i": "2006-01-02 15:04",
		"s": "2006-01-02 15:04:05",
	}
)

// StrToUnixTime 字符串转时间戳
func StrToUnixTime(str string) (int64, error) {
	layout := "2006-01-02 15:04:05"
	t, err := time.ParseInLocation(layout, str, time.Local)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

func StrToTime(str string) (time.Time, error) {
	layout := "2006-01-02 15:04:05"
	return time.ParseInLocation(layout, str, time.Local)
}
func StrToTime1(str string) (time.Time, error) {
	layout := "2006-01-02"
	return time.ParseInLocation(layout, str, time.Local)
}

// TimeToStr 时间戳转字符串
func TimeToStr(timer int64) string {
	tm := time.Unix(timer, 0)
	return tm.Format("2006-01-02 15:04:05")
}

func GetAnyDayStartAndEndTime(dateNow time.Time) (startTime, endTime time.Time) {
	startTime = time.Date(dateNow.Year(), dateNow.Month(), dateNow.Day(), 0, 0, 0, 0, dateNow.Location())
	endTime = time.Date(dateNow.Year(), dateNow.Month(), dateNow.Day(), 23, 59, 59, 0, dateNow.Location())
	return
}

// GetDayHourTime 获取一天的整点时间
func GetDayHourTime(hour int) time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), hour, 0, 0, 0, now.Location())
}

// GetNextDayHourTime 获取明天的整点时间
func GetNextDayHourTime(hour int) time.Time {
	nextDay := time.Now().Add(time.Hour * 24)
	return time.Date(nextDay.Year(), nextDay.Month(), nextDay.Day(), hour, 0, 0, 0, nextDay.Location())
}

// GetLastDayHourTime 获取昨天的整点之间
func GetLastDayHourTime(hour int) time.Time {
	lastDay := time.Now().Add(-1 * time.Hour * 24)
	return time.Date(lastDay.Year(), lastDay.Month(), lastDay.Day(), hour, 0, 0, 0, lastDay.Location())
}

// GetDayStartTime 获取一天的开始时间
func GetDayStartTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// GetDayEndTime 获取一天的结束
func GetDayEndTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
}

// GetToDayDateOfMonth 获取今天是一个月内的第几天 比如4月23号就是第23天
func GetToDayDateOfMonth() int {
	return time.Now().Day()
}

// FormatTimeToDayStr 把日期转换成天，格式 2006-01-02
func FormatTimeToDayStr(t time.Time) string {
	return t.Format("2006-01-02")
}

// FormatTimeToMonthDayStr 把日期转换成天，格式 01-02
func FormatTimeToMonthDayStr(t time.Time) string {
	return t.Format("01-02")
}

// GetFormDateTimeTime 获取具体时间的time.Time
func GetFormDateTimeTime(timeString string, unit string) (timeToTime time.Time) {
	loc, _ := time.LoadLocation("Local")

	layout, ok := TimeLayoutMap[unit]
	if !ok {
		layout = TimeLayoutMap["s"]
	}

	timeToTime, _ = time.ParseInLocation(layout, timeString, loc)

	return
}

// GetCurrentTimeOffsetTime 获取具体时间偏移一定天数后的时间
func GetCurrentTimeOffsetTime(toDayTime string, unit string, offset int) (offsetTime string) {
	loc, _ := time.LoadLocation("Local")
	layout, ok := TimeLayoutMap[unit]
	if !ok {
		layout = TimeLayoutMap["s"]
	}
	res, _ := time.ParseInLocation(layout, toDayTime, loc)

	offsetTimeStamp := time.Date(res.Year(), res.Month(), res.Day(), res.Hour(), res.Minute(), res.Second(), 0, time.Local).AddDate(0, 0, offset)
	offsetTime = offsetTimeStamp.Format("2006-01-02 15:04:05")

	return offsetTime
}

// GetTwoTimeSalt 获取两个时间之间的差值: d:天\h:小时\
func GetTwoTimeSalt(firstTime string, secondTime string, unit string, offsetUnit string) (salt float64) {
	// 获取第一个参数time.Time
	firstTimeTime := GetFormDateTimeTime(firstTime, unit)
	// 获取第二个参数的time.Time
	secondTimeTime := GetFormDateTimeTime(secondTime, unit)

	// 计算时间差
	duration := firstTimeTime.Sub(secondTimeTime)

	// 返回差
	if offsetUnit == "d" {
		salt = duration.Minutes() / 60 / 24
	}
	if offsetUnit == "h" {
		salt = duration.Minutes() / 60
	}
	if offsetUnit == "s" {
		salt = duration.Seconds()
	}

	return salt
}

// GetLaterTime 获取传入时间n天后的时间
func GetLaterTime(oldTime string, num int) (newTime string) {
	loc, _ := time.LoadLocation("Local")
	oldTimeStamp, _ := time.ParseInLocation("2006-01-02 15:04:05", oldTime, loc)
	newTimeStamp := oldTimeStamp.AddDate(0, 0, num)

	return newTimeStamp.Format("2006-01-02")
}

func LaterTime(oldTime string, num int) (newTime time.Time) {
	loc, _ := time.LoadLocation("Local")
	oldTimeStamp, _ := time.ParseInLocation("2006-01-02 15:04:05", oldTime, loc)
	newTimeStamp := oldTimeStamp.AddDate(0, 0, num)
	res, _ := time.ParseInLocation("2006-01-02 15:04:05", newTimeStamp.Format("2006-01-02 15:04:05"), loc)

	return res
}

// TimeStrExchangeTimeStampV2 具体时间转时间戳
func TimeStrExchangeTimeStampV2(timeStr string, unit string) (response int64, err error) {
	loc, _ := time.LoadLocation("Local")
	layout, ok := TimeLayoutMap[unit]
	if !ok {
		layout = TimeLayoutMap["s"]
	}
	res, err := time.ParseInLocation(layout, timeStr, loc)
	response = res.Unix()

	return
}

func GetMonthStartAndEndDay(year, month int) (firstOfMonth, lastOfMonth time.Time) {
	firstOfMonth = time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	lastOfMonth = firstOfMonth.AddDate(0, 1, 0).Add(time.Second * -1)
	return
}

// GetExcelTimeFormat excel 时间标准化
func GetExcelTimeFormat(time1 string) (Time time.Time, err error) {
	timeMoth := ""
	timeDay := ""
	timeStrlList := strings.Split(time1, "/")
	if len(timeStrlList[1]) < 2 {
		timeMoth = "0" + timeStrlList[1]
	} else {
		timeMoth = timeStrlList[1]
	}
	if len(timeStrlList[2]) < 2 {
		timeDay = "0" + timeStrlList[2]
	} else {
		timeDay = timeStrlList[2]
	}
	newTime1 := timeStrlList[0] + "/" + timeMoth + "/" + timeDay
	time3 := ""
	time2 := strings.Replace(cast.ToString(newTime1), "/", "-", 2)
	str1List := strings.Split(cast.ToString(time2), " ")
	if len(str1List) == 2 {
		str2List := strings.Split(cast.ToString(str1List[1]), ":")
		if len(str2List) == 2 {
			time3 = time2 + ":00"
		} else if len(str2List) == 3 {
			time3 = time2
		}
	}
	// time2 := strings.Replace(cast.ToString(time1), "/", "-", 2) + ":00"
	Time, err = time.ParseInLocation(TimeLayoutMap["s"], time3, time.Local) // 这里按照当前时区转
	if err != nil {
		return
	}
	return
}

// GetExcelDateTimeFormat Excel 年月日的时间
func GetExcelDateTimeFormat(time1 string) (Time time.Time, err error) {
	timeMoth := ""
	timeDay := ""
	timeStrlList := strings.Split(time1, "/")
	if len(timeStrlList[1]) < 2 {
		timeMoth = "0" + timeStrlList[1]
	} else {
		timeMoth = timeStrlList[1]
	}
	if len(timeStrlList[2]) < 2 {
		timeDay = "0" + timeStrlList[2]
	} else {
		timeDay = timeStrlList[2]
	}
	newTime1 := timeStrlList[0] + "/" + timeMoth + "/" + timeDay
	time2 := strings.Replace(cast.ToString(newTime1), "/", "-", 2)
	// time2 := strings.Replace(cast.ToString(time1), "/", "-", 2) + ":00"
	Time, err = time.ParseInLocation(TimeLayoutMap["d"], time2, time.Local) // 这里按照当前时区转
	if err != nil {
		return
	}
	return
}

func GetLaterDaterTime(oldTime string, num int) (newTime string) {
	loc, _ := time.LoadLocation("Local")
	oldTimeStamp, _ := time.ParseInLocation("2006-01-02", oldTime, loc)
	newTimeStamp := oldTimeStamp.AddDate(0, 0, num)

	return newTimeStamp.Format("2006-01-02")
}

// TimeStampToYearMonth 将时间戳转换为对应所处的年月 例如：2023-10
func TimeStampToYearMonth(timestamp int64) (newTime string) {
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01")
}

// TimeStampToYearMonth2 将时间戳转换为对应所处的年月 例如：2023年10月
func TimeStampToYearMonth2(timestamp int64) (newTime string) {
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006年1月")
}

// TimeStampToYearMonthDay 将时间转换为对应所处的年月日 例如：2023年10月1日
func TimeStampToYearMonthDay(time time.Time) (newTime string) {
	return time.Format("2006年1月2日")
}

// TimeStampToDateStr 将时间戳转换为对应时间字符串展示 例如：2006-01-02 15:04:05
func TimeStampToDateStr(timestamp int64) (newTime string) {
	if timestamp == 0 {
		return ""
	}
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 15:04:05")
}

// TimeStampBeforeMonths 获取当前时间对应的指定months个月之前的这个月的第一天的起始时间戳
/*
比如今天是2024年9月12日，months = 24 res = 2022年9第一天的开始时间戳 2022-09-01 00:00:00
*/
func TimeStampBeforeMonths(months int) (newTime int64) {
	now := GetMonthStartDayTime(time.Now())
	past := now.AddDate(0, -months, 0)
	pastFirstOfMonth := time.Date(past.Year(), past.Month(), 1, 0, 0, 0, 0, now.Location())
	return pastFirstOfMonth.Unix()
}

func GetMonthStartDay(monthTime time.Time) string {
	dayTime := time.Date(monthTime.Year(), monthTime.Month(), 1, 0, 0, 0, 0, time.Local)
	newTimeStr := TimeStampToYearMonthDay(dayTime)
	return newTimeStr
}

func GetMonthStartDayTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local)
}

func GetNextMonthStartTime(t time.Time) time.Time {
	return GetMonthStartDayTime(t).AddDate(0, 1, 0)
}

// GetIntervalTime 获取当前到指定时间的间隔，单位是time.Duration
func GetIntervalTime(t time.Time) time.Duration {
	interval := t.Sub(time.Now())
	// 如何指定时间小于当前时间，返回1s间隔
	if interval <= 0 {
		return time.Second
	}
	return interval
}

// GetWeekStartAndEnd 获取本周的开始时间和结束时间
func GetWeekStartAndEnd() (time.Time, time.Time) {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	monday := now.AddDate(0, 0, offset)
	weekStart := time.Date(monday.Year(), monday.Month(), monday.Day(), 0, 0, 0, 0, time.Local)
	// 本周结束时间（周日）
	weekEnd := weekStart.AddDate(0, 0, 6).Add(time.Hour*23 + time.Minute*59 + time.Second*59)

	return weekStart, weekEnd
}

// GetMonday 获取周一，想要什么格式自己传
func GetMonday(layout string) string {
	weekStart, _ := GetWeekStartAndEnd()

	if layout == "" {
		layout = "2006-01-02"
	}

	return weekStart.Format(layout)
}

// GetTwoDateIntervalDuration 获取开始时间到结束时间的时间间隔
func GetTwoDateIntervalDuration(startTime time.Time, endTime time.Time) time.Duration {
	return endTime.Sub(startTime)
}
