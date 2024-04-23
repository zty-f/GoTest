package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"
)

func main() {
	ids := []int{2000220581, 2000220582}
	for _, id := range ids {
		grant(id, 10000)
	}
}

func grant(stuId, growthValue int) {
	url := "https://xwadmin.xiwang.com/growthcenteradmin/v1/admin/growthValue/grant"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("stuId", strconv.Itoa(stuId))
	_ = writer.WriteField("growthValue", strconv.Itoa(growthValue))
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("cookie", "xesId=93d14e532aad87b7174def9fe45e5c1e; is_login=1; xes_acc=eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJ1aWQiOiIyMTAwMDAyMDA1IiwidXNlcl9pZCI6IjIxMDAwNTE2ODQiLCJyb2xlIjoxLCJleHAiOjE3MDk4ODk4NTksInZlciI6IjEuMCJ9.NYu7CC5vb2ep3RmDFbCYRxNq2Ii9dPCOsXZRV-IBjDDUuaSIdlfqBBxMhVPX9VUStFck_mBl6ei83u_93Rb7K548GSzAjmcPiioqMK3vwVabOf8tTsJ6XTYL3NTc53pV2cN_LUdRhS7zSUn4xlKVrsBl1XQPcCvIIn9uawYudZm1Up_60VnlFU6aFu5ameoJU7qg1qWE36AVAi-gOlY8bvdxd6JVdwiYxdmV0zvwh9Ayp9ZH8eFvjAxrDqAGFTe0VXeA0D7ww7-Lkdo1BmciUMTacenxAy52hHkisCTn8E4VJ3k_pelR--0HEbv6fSdOJ0khX4kyPnQPwrBuguBnlw; userGrade=2; stu_id=2100051684; stu_name=syhtest2100051684%40xiwang.com; tfstk=fcFvyFmZoaLY2WZFKR6uQAR6ES_lE5UqeozBscmDCuE8uyemCC2DffE8YANmirq-5uaLil4bjlKTmuqtQZoifcEg5-blt644ufl_XMfhtgM5zCExlqMf50gElEFCHL44uf8PfTg4QrJtTQne6fZsNbgmWfgscli5N0uZ5ng6GaUWWxTR5gWDgJFol9ShuF-awz0xHcdIqCdCSVkxArMPAMNR1xnQl0d1vaXUm0gLNG-mb7U7GWE5ZIhTPVEKG8IWMlUbLluQc_Kjk-aL-YFNAC3zFokok8Q6FmaIcyG0s3_xbSqY6xPCxQobe7Pti8jvtck14TV39bVDx4lJhaQJQdkjzKfPRyyeQVLZyDbordJZHG3-xaQJQdkjz4nhr4vwQxIO.; token=eyJhcHBpZCI6IjIzNSIsInRva2VuIjoiMGVkNGU5NThhZmY3NzVhM2I4NDg1NjE5ODc5ZDdjNmMifQ%3D%3D; MKC-ExternalAccessToken-3=eyJhcHBpZCI6IjIzNSIsInRva2VuIjoiMGVkNGU5NThhZmY3NzVhM2I4NDg1NjE5ODc5ZDdjNmMifQ==; currentGrade=2; stu_area_id=100; isVisitFirst=1; wx=0009af0eea3f9d3ba8eb27f077f1b85f4xxqxsfy0; prelogid=7884fddecabf447641fe05eb799eb8e3; TM-SessionID=5c22465afbdb2f72bdb29f9259129697; tal_token=tal1731wGQeiQP7W6FDNyhfXhBYWq1a4iwMflsTmbmCW6aiiP5qVCckN-B70otvCsa0DAC3mdb8lfgqfJ8zkqHviHRTHQROZVsqoY_FM9zRQDh4C-bUbRaBS0hJJ33LMkmqvOO2owDIBk4zfoWgQ2GZY4aUoG7F9yDMm4OaNpDp2pEqSEmko_XNbVLDh7Gd1l2nKGLXo9EEg9ClV7rQxXC4bQ0UqRClYIILAmlSC-yXh6mcxg-Cg4ahVqvnwlcO7GNq1FCKmj4KAC1vptQ-1cu9TNQ6hl6Az638WQrIgrelJmnF35e5iVJpxYkG-oCf6BYpPw_ubQ")
	req.Header.Add("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoi5byg5aSp5rOzIiwid29ya19jb2RlIjoiVzAwODQ5MiIsImF2YXRhciI6Imh0dHBzOi8veWFjaC1zdGF0aWMuemhpeWlubG91LmNvbS9vbmxpbmUvanNhcGkvMTcxMDcyNzQ5NTM0My9rZTNmNThrbGRqOC9hNzY3M2JmNC1kYTk2LTQ0ZTAtOGYxMC05NmU3OWZjNjI4NjgucG5nIiwiZW1haWwiOiJ3X3poYW5ndGlhbnlvbmdAeGl3YW5nLmNvbSIsImV4cCI6MTcxMzk1MzgzNywibmJmIjoxNzEzODY3NDM2fQ.JzVeCuGXugZ0MP_p9ReDbctgCJ_QQ4JVxXtiIxqjZo4")
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://apifox.com)")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Host", "xwadmin.xiwang.com")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Content-Type", "multipart/form-data; boundary=--------------------------414124545713488125546066")

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
