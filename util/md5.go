package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
)

func Md51(str string) string {
	data := []byte(str)
	md5New := md5.New()
	md5New.Write(data)
	// hex转字符串
	md5String := hex.EncodeToString(md5New.Sum(nil))
	return md5String
}

func Md52(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	sum := h.Sum(nil)
	return hex.EncodeToString(sum)
}

func Md53(str string) string {
	data := []byte(str)
	sum := md5.Sum(data)
	// hex转字符串
	md5String := hex.EncodeToString(sum[:])
	return md5String
}

func Md54(str string) string {
	data := []byte(str)
	md5New := md5.New()
	md5New.Write(data)
	md5Sum := md5New.Sum(nil)
	return fmt.Sprintf("%x", md5Sum)
}

func Md55(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	md5Sum := h.Sum(nil)
	return fmt.Sprintf("%x", md5Sum)
}

func Md56(str string) string {
	data := []byte(str)
	sum := md5.Sum(data)
	return fmt.Sprintf("%x", sum)
}
