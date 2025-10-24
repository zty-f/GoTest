package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// APIClient API客户端
type APIClient struct {
	BaseURL string
	Cookie  string
	Client  *http.Client
}

// NewAPIClient 创建新的API客户端
func NewAPIClient(baseURL, cookie string) *APIClient {
	return &APIClient{
		BaseURL: baseURL,
		Cookie:  cookie,
		Client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// GetBookData 获取绘本数据
func (c *APIClient) GetBookData(bookID int64) (*APIResponse, error) {
	// 构建请求体
	requestBody := map[string]interface{}{
		"bookid": bookID,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("序列化请求体失败: %v", err)
	}

	// 创建请求
	req, err := http.NewRequest("POST", c.BaseURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}

	// 设置请求头
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Origin", "https://sea.pri.ibanyu.com")
	req.Header.Set("Priority", "u=1, i")
	req.Header.Set("Referer", "https://sea.pri.ibanyu.com/character-admin/")
	req.Header.Set("Sec-Ch-Ua", `"Google Chrome";v="141", "Not?A_Brand";v="8", "Chromium";v="141"`)
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", `"macOS"`)
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36")

	if c.Cookie != "" {
		req.Header.Set("Cookie", c.Cookie)
	}

	// 发送请求
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("发送请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 检查响应状态
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("请求失败，状态码: %d", resp.StatusCode)
	}

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应体失败: %v", err)
	}

	// 解析JSON
	var apiResponse APIResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return nil, fmt.Errorf("解析JSON失败: %v", err)
	}

	// 检查API返回码
	if apiResponse.Ret != 1 || apiResponse.Code != 0 {
		return nil, fmt.Errorf("API返回错误: ret=%d, code=%d", apiResponse.Ret, apiResponse.Code)
	}

	return &apiResponse, nil
}
