package gemini

import (
	"fmt"
	"github.com/google/generative-ai-go/genai"
	"moechat/core/api/part"
)

func transformToProviderMessages(msgs []part.Message) ([]*genai.Content, part.Message, error) {
	if len(msgs) == 0 {
		return nil, *new(part.Message), fmt.Errorf("no messages to transform")
	}

	var genaiMessages []*genai.Content

	for _, msg := range msgs[:len(msgs)-1] {
		genaiMessages = append(genaiMessages, &genai.Content{
			Role: msg.Role,
			Parts: []genai.Part{
				//genai.FileData{URI: file.URI},
				genai.Text(msg.Content),
			},
		})
	}

	// Process the last message
	lastMsg := msgs[len(msgs)-1]

	return genaiMessages, lastMsg, nil
}
