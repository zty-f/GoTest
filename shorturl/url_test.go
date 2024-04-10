package main

import (
	"fmt"
	"net/url"
	"testing"
)

func TestParse(t *testing.T) {
	SourceUrl := "https%3A%2F%2Fldx-test.wen-su.com%2Fcps-test%2Findex.html%3FA614D4C9318D2C729186B987EF2C13B43D7C33B8C6D3CA65C19CEF009A063F87D5A9C5328609D9DB006111109648A05BF0294C890D18886E5AED11C7A9CAD4968EF180945CD143E8AF4AA68093AF91D24715CB50C92394B79D238BD0796B84626A676A72558910D17F7E54199309A0B7A6ECDDF1D31B808D1E1DA268DD14621D"
	sourceDecodeUrl, err := url.QueryUnescape(SourceUrl) //解码
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(sourceDecodeUrl)

	// 校验域名是否白名单
	parsedURL, err := url.Parse(sourceDecodeUrl)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(parsedURL)
}
