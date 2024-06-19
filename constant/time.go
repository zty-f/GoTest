package constant

import "time"

const (
	TenMinute  = time.Minute * 10
	TenSecond  = time.Second * 10
	Day        = 24 * time.Hour
	SevenDay   = 7 * Day
	EightDay   = 8 * Day
	Month      = 30 * Day
	TwoMonth   = 2 * Month
	ThreeMonth = 3 * Month
	SixMonth   = 6 * Month
	Year       = 365 * Day

	MaxRetrySleepDuration = time.Minute //最大重试休眠时间

	WaitKafkaCommitMessageDuration = time.Second * 10 //等待kafka消息提交成功时间

	NinetyDay = 90 * Day

	PanicTickerDuration              = time.Minute     //打印panic日志间隔
	FinishRealTimeTaskLockExpireTime = time.Second * 3 //完成实时任务锁过期时间
	AddUserMedalCacheLockExpireTime  = time.Second * 3 //添加用户奖牌缓存防并发锁过期时间
)
