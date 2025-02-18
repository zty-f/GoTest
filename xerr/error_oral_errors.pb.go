// Code generated by gen-go-errors. DO NOT EDIT.

package xerr

import (
	fmt "fmt"
	errors "test/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

// biz error
// 系统:oral 12 口算
// 模块:通用 00
func IsOralPhotoOcrJudgementOverLimit(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Code == 1200000
}

// biz error
// 系统:oral 12 口算
// 模块:通用 00
func ErrorOralPhotoOcrJudgementOverLimit(format string, args ...interface{}) *errors.Error {
	if format == "" {
		return errors.New(errors.Stat_FAILED, 1200000, "今天已经检查好多啦 休息一下明天再来吧～")
	}
	return errors.New(errors.Stat_FAILED, 1200000, fmt.Sprintf(format, args...))
}

func IsOralPhotoOcrJudgementRequestTooFast(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Code == 1200001
}

func ErrorOralPhotoOcrJudgementRequestTooFast(format string, args ...interface{}) *errors.Error {
	if format == "" {
		return errors.New(errors.Stat_FAILED, 1200001, "请求速度过快")
	}
	return errors.New(errors.Stat_FAILED, 1200001, fmt.Sprintf(format, args...))
}

func IsOralCommon(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Code == 1200002
}

func ErrorOralCommon(format string, args ...interface{}) *errors.Error {
	if format == "" {
		return errors.New(errors.Stat_FAILED, 1200002, "哎呀，好像出现了一点问题")
	}
	return errors.New(errors.Stat_FAILED, 1200002, fmt.Sprintf(format, args...))
}

func IsOralPhotoOcrJudgementImgTooLarge(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Code == 1200003
}

func ErrorOralPhotoOcrJudgementImgTooLarge(format string, args ...interface{}) *errors.Error {
	if format == "" {
		return errors.New(errors.Stat_FAILED, 1200003, "图片尺寸过大")
	}
	return errors.New(errors.Stat_FAILED, 1200003, fmt.Sprintf(format, args...))
}

func IsOralPhotoOcrJudgementImgTypeErr(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Code == 1200004
}

func ErrorOralPhotoOcrJudgementImgTypeErr(format string, args ...interface{}) *errors.Error {
	if format == "" {
		return errors.New(errors.Stat_FAILED, 1200004, "图片格式错误")
	}
	return errors.New(errors.Stat_FAILED, 1200004, fmt.Sprintf(format, args...))
}
