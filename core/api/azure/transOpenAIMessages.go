package azure

import (
	"encoding/json"
	"fmt"
	"github.com/sashabaranov/go-openai"
)

func transformToProviderMessages(msgs json.RawMessage) ([]openai.ChatCompletionMessage, error) {
	// 将 json.RawMessage 解析为 []map[string]interface{}
	var messages []map[string]any
	if err := json.Unmarshal(msgs, &messages); err != nil {
		return nil, fmt.Errorf("failed to unmarshal messages: %w", err)
	}

	var openAIMessages []openai.ChatCompletionMessage
	for _, msg := range messages {
		role, ok := msg["role"].(string)
		if !ok {
			return nil, fmt.Errorf("invalid role type, expected string")
		}
		content, ok := msg["content"].(string)
		if !ok {
			return nil, fmt.Errorf("invalid content type, expected string")
		}

		openAIMessages = append(openAIMessages, openai.ChatCompletionMessage{
			Role:    role,
			Content: content,
		})
	}
	return openAIMessages, nil
}
