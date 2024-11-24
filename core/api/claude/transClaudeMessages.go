package claude

import (
	"github.com/liushuangls/go-anthropic/v2"
	"moechat/core/api/part"
)

func transformToProviderMessages(msgs []part.Message) ([]anthropic.Message, error) {
	var claudeMessages []anthropic.Message
	for _, msg := range msgs {
		if msg.Role == "user" {
			claudeMessages = append(claudeMessages, anthropic.NewUserTextMessage(msg.Content))
		} else if msg.Role == "assistant" {
			claudeMessages = append(claudeMessages, anthropic.NewAssistantTextMessage(msg.Content))
		}
	}

	return claudeMessages, nil
}
