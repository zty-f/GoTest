package dify

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

const (
	// ResponseModeStreaming 流式响应模式
	ResponseModeStreaming = "streaming"
	// ResponseModeBlocking 阻塞响应模式
	ResponseModeBlocking = "blocking"
)

// 单例实例和锁
var (
	qwAIChatClient *QwAIChatClient
	qwAIChatOnce   sync.Once
)

// QwAIChatClient Qw AI 客户端
type QwAIChatClient struct {
	config     Config
	httpClient *http.Client
}

// Config 配置结构体
type Config struct {
	APIKey       string `json:"api_key"`
	BaseURL      string `json:"base_url"`
	ResponseMode string `json:"response_mode"` // streaming or blocking
}

// FileInput 文件输入结构体
type FileInput struct {
	Type           string `json:"type"`                     // image
	TransferMethod string `json:"transfer_method"`          // remote_url or local_file
	URL            string `json:"url,omitempty"`            // 仅当 transfer_method 为 remote_url 时
	UploadFileID   string `json:"upload_file_id,omitempty"` // 仅当 transfer_method 为 local_file 时
}

// NewQwAIChatClient 创建新的 Workflow 客户端
func NewQwAIChatClient(responseMode string) *QwAIChatClient {
	qwAIChatOnce.Do(func() {
		if qwAIChatClient == nil {
			qwAIChatClient = &QwAIChatClient{
				config: Config{
					APIKey:       "app-jRtqGS59oaV5bBKetMPercQd",
					BaseURL:      "http://10.120.194.12/v1",
					ResponseMode: responseMode,
				},
				httpClient: &http.Client{
					Timeout: 120 * time.Second, // 设置超时时间
				},
			}
		}
	})
	return qwAIChatClient
}

type RunQwAIChatRequest struct {
	Query            string                 `json:"query"`                        // 用户输入
	ConversationId   string                 `json:"conversation_id"`              // 会话 ID
	Inputs           map[string]interface{} `json:"inputs"`                       // 输入变量
	ResponseMode     string                 `json:"response_mode"`                // streaming or blocking
	User             string                 `json:"user"`                         // 用户标识
	Files            []FileInput            `json:"files,omitempty"`              // 文件列表
	AutoGenerateName bool                   `json:"auto_generate_name,omitempty"` // （选填）自动生成标题，默认 true。 若设置为 false，则可通过调用会话重命名接口并设置 auto_generate 为 true 实现异步生成标题。
}

func (c *QwAIChatClient) QwAIChat(ctx context.Context, inputs map[string]interface{}, query, conversationId, user string) (*ChatCompletionResponse, error) {
	// fun := "QwAIChatClient.QwAIChat"
	result, err := c.RunQwAIChat(ctx, inputs, query, conversationId, user)
	if err != nil {
		return nil, err
	}
	if response, ok := result.(*ChatCompletionResponse); ok {
		if response != nil && response.MessageId == "" {
			return nil, fmt.Errorf("QwAIChat execution failed: %s", response.Answer)
		}
		return response, nil
	}
	return nil, fmt.Errorf("unexpected response type: %T", result)
}

// RunQwAIChat 执行 Qw AI 聊天
func (c *QwAIChatClient) RunQwAIChat(ctx context.Context, inputs map[string]interface{}, query, conversationId, user string) (interface{}, error) {
	request := RunQwAIChatRequest{
		Query:          query,
		ConversationId: conversationId,
		User:           user,
		ResponseMode:   c.config.ResponseMode,
		Inputs:         inputs,
	}

	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
		fmt.Sprintf("%s/chat-messages", c.config.BaseURL),
		bytes.NewBuffer(requestBody),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.config.APIKey))
	req.Header.Set("Content-Type", "application/json")

	if c.config.ResponseMode == ResponseModeBlocking {
		return c.handleBlockingResponse(ctx, req)
	}
	return c.handleStreamingResponse(ctx, req)
}

// ChatCompletionResponse 阻塞模式响应结构体
type ChatCompletionResponse struct {
	MessageId      string `json:"message_id"`
	ConversationId string `json:"conversation_id"`
	Mod            string `json:"mod"`
	Answer         string `json:"answer"`
	MetaData       any    `json:"metadata"`
	CreatedAt      int64  `json:"created_at"`
}

// handleBlockingResponse 处理阻塞模式响应
func (c *QwAIChatClient) handleBlockingResponse(ctx context.Context, req *http.Request) (*ChatCompletionResponse, error) {
	// fun := "QwAIChatClient.handleBlockingResponse"
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		fmt.Sprintf("QwAIChatClient.handleBlockingResponse failed: %s", string(body))
		return nil, errors.New(string(body))
	}

	fmt.Sprintf("QwAIChatClient.handleBlockingResponse success: %s", string(body))

	var response ChatCompletionResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// ChatStreamEvent 流式事件通用结构
type ChatStreamEvent struct {
	Event          string `json:"event"`
	TaskID         string `json:"task_id"`
	MessageId      string `json:"message_id"`
	ConversationId string `json:"conversation_id"`
	Answer         string `json:"answer"`
	MetaData       any    `json:"metadata"`
	CreateAt       int64  `json:"create_at"`
}

// handleStreamingResponse 处理流式响应
func (c *QwAIChatClient) handleStreamingResponse(ctx context.Context, req *http.Request) (chan ChatStreamEvent, error) {
	_ = "QwAIChatClient.handleStreamingResponse"
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		_, _ = io.ReadAll(resp.Body)
		resp.Body.Close()
		return nil, err
	}

	// 创建事件通道
	eventChan := make(chan ChatStreamEvent)

	// 启动 goroutine 处理流式响应
	go func() {
		defer resp.Body.Close()
		defer close(eventChan)

		decoder := json.NewDecoder(resp.Body)
		for {
			var event ChatStreamEvent
			if err := decoder.Decode(&event); err != nil {
				if errors.Is(err, io.EOF) {
					return
				}
				return
			}
			eventChan <- event
		}
	}()

	return eventChan, nil
}
