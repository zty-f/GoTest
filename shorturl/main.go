package main

import (
	"crypto/md5"
	"fmt"
	"math"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func main() {
	//shortUrl := "OWJdbu"
	//sourceUrl := "www.baidu.com"
	//getSign(sourceUrl)
	//hashValue := DecodeStr(shortUrl)
	//newTable := getTable(uint(hashValue))
	//fmt.Println(newTable)
	url1, _ := url.QueryUnescape("%09https%3A%2F%2Fac.xiwang.com%2Fcampus-activity%2F%23%2F%3Fuuid%3D42a0a066718542809cb74a1ccfa0ab56")
	fmt.Println(url1)
	//url1 = strings.TrimSpace(url1)
	fmt.Println(url1)
	parsedURL, err := url.Parse(url1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(parsedURL)
}

// DecodeStr 62进制到int进制
func DecodeStr(str string) int {
	const CHARS = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var num int
	n := len(str)
	for i := 0; i < n; i++ {
		pos := strings.IndexByte(CHARS, str[i])
		num += int(math.Pow(62, float64(n-i-1)) * float64(pos))
	}
	return num
}

func getTable(hashValue uint) (TableName string) {
	j := hashValue % 32
	TableName = "short_url_" + strconv.Itoa(int(j))
	return TableName
}

func getSign(sourceUrl string) {
	key := "jXH33mAf"
	unix := time.Now().Unix()
	fmt.Println("当前时间戳：", unix)
	sign := md5.Sum([]byte(sourceUrl + strconv.Itoa(int(unix)) + key))
	fmt.Println("sign：", fmt.Sprintf("%x", sign))
}
