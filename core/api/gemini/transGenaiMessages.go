package gemini

import (
	"encoding/json"
	"fmt"
	"github.com/google/generative-ai-go/genai"
)

func transformToProviderMessages(msgs json.RawMessage) ([]*genai.Content, map[string]any, error) {
	var messages []map[string]any
	if err := json.Unmarshal(msgs, &messages); err != nil {
		return nil, nil, fmt.Errorf("failed to unmarshal messages: %w", err)
	}

	if len(messages) == 0 {
		return nil, nil, fmt.Errorf("no messages to transform")
	}

	var genaiMessages []*genai.Content
	// Process all except the last message
	for _, msg := range messages[:len(messages)-1] {
		role, ok := msg["role"].(string)
		if !ok {
			return nil, nil, fmt.Errorf("invalid role type, expected string")
		}
		content, ok := msg["content"].(string)
		if !ok {
			return nil, nil, fmt.Errorf("invalid content type, expected string")
		}
		if role == "assistant" {
			role = "model"
		}
		genaiMessages = append(genaiMessages, &genai.Content{
			Role: role,
			Parts: []genai.Part{
				//genai.FileData{URI: file.URI},
				genai.Text(content),
			},
		})
	}

	// Process the last message
	lastMsg := messages[len(messages)-1]

	return genaiMessages, lastMsg, nil
}
