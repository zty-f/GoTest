package dify

import (
	"context"
	"testing"
)

func TestQwAIChatClient_RunQwAIChat(t *testing.T) {
	client := NewQwAIChatClient(ResponseModeBlocking)
	ctx := context.Background()
	inputs := map[string]interface{}{
		"user_text":  "学习多久合适",
		"user_id":    "3290014997",
		"student_id": "3290014997",
	}
	response, err := client.QwAIChat(ctx, inputs, "需要多少时间", "", "3290014997")
	if err != nil {
		t.Fatalf("RunQwAIChat failed: %v", err)
	}
	t.Log(response)
}
