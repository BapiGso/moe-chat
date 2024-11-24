package claude

import (
	"encoding/json"
	"fmt"
	"github.com/liushuangls/go-anthropic/v2"
)

func transformToProviderMessages(msgs json.RawMessage) ([]anthropic.Message, error) {

	var messages []map[string]any
	if err := json.Unmarshal(msgs, &messages); err != nil {
		return nil, fmt.Errorf("failed to unmarshal messages: %w", err)
	}

	var claudeMessages []anthropic.Message
	for _, msg := range messages {
		role, ok := msg["role"].(string)
		if !ok {
			return nil, fmt.Errorf("invalid role type, expected string")
		}
		content, ok := msg["content"].(string)
		if !ok {
			return nil, fmt.Errorf("invalid content type, expected string")
		}
		if role == "user" {
			claudeMessages = append(claudeMessages, anthropic.NewUserTextMessage(content))
		} else if role == "assistant" {
			claudeMessages = append(claudeMessages, anthropic.NewAssistantTextMessage(content))
		}

	}
	return claudeMessages, nil
}
