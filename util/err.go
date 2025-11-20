package util

import "strings"

// ErrorsIsDuplicate 判断错误是否是数据重复错误
func ErrorsIsDuplicate(err error) bool {
	if err == nil {
		return false
	}

	if strings.Index(strings.ToLower(err.Error()), "duplicate") != -1 {
		return true
	}

	return false
}

// ErrorsIsLevelNoMatch 判断错误是否是等级不符合当前操作
func ErrorsIsLevelNoMatch(err error) bool {
	if err == nil {
		return false
	}

	if strings.Index(strings.ToLower(err.Error()), "等级及以上可使用") != -1 {
		return true
	}

	return false
}

// ErrorsIsUnLock 判断错误是否是没有解锁
func ErrorsIsUnLock(err error) bool {
	if err == nil {
		return false
	}

	if strings.Index(strings.ToLower(err.Error()), "未解锁") != -1 {
		return true
	}

	return false
}

// ErrorsIsChecking 判断错误是审核中
func ErrorsIsChecking(err error) bool {
	if err == nil {
		return false
	}

	if strings.Index(strings.ToLower(err.Error()), "审核中") != -1 {
		return true
	}

	return false
}

// ErrorsIsEOF 判断错误是否是EOF相关
func ErrorsIsEOF(err error) bool {
	if err == nil {
		return false
	}

	if strings.Index(strings.ToUpper(err.Error()), "EOF") != -1 {
		return true
	}

	return false
}

// ErrorsIsNoPrize 判断错误是否是没有奖品
func ErrorsIsNoPrize(err error) bool {
	if err == nil {
		return false
	}

	if strings.Index(strings.ToLower(err.Error()), "no prize") != -1 {
		return true
	}

	return false
}
