package gemini

import (
	"encoding/base64"
	"fmt"
	"github.com/google/generative-ai-go/genai"
	"github.com/labstack/echo/v4"
	"moechat/core/api/part"
	"moechat/core/database"
)

func transformToProviderMessages(ctx echo.Context, msgs []part.Message) ([]*genai.Content, part.Message, error) {
	if len(msgs) == 0 {
		return nil, *new(part.Message), fmt.Errorf("no messages to transform")
	}

	var genaiMessages []*genai.Content

	for _, msg := range msgs[:len(msgs)-1] {
		if msg.Files != nil {
			var file database.File
			for _, f := range msg.Files {
				if err := database.DB.Get(&file,
					`SELECT * from file WHERE hash = ? AND email = ?`, f.Hash, ctx.Get("email")); err != nil {
					return nil, *new(part.Message), err
				}
				genaiMessages = append(genaiMessages, &genai.Content{
					Role: msg.Role,
					Parts: []genai.Part{
						genai.FileData{
							MIMEType: file.MimeType,
							URI:      base64.StdEncoding.EncodeToString(file.Data),
						},
					},
				})
			}
		}
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
