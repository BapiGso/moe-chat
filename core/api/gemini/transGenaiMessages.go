package gemini

import (
	"github.com/google/generative-ai-go/genai"
	"github.com/labstack/echo/v4"
	"moechat/core/api/part"
	"moechat/core/database"
)

func transformToProviderMessages(ctx echo.Context, cs *genai.ChatSession, msgs []part.Message) (*genai.Content, error) {
	if len(msgs) == 1 {
		return &genai.Content{
			Role: msgs[0].Role,
			Parts: []genai.Part{
				genai.Text(msgs[0].Content),
			},
		}, nil
	}
	for _, msg := range msgs[:len(msgs)-1] {
		if msg.Files != nil {
			var file database.File
			var tmpparts []genai.Part
			for _, f := range msg.Files {
				err := database.DB.Get(&file,
					`SELECT * from file WHERE hash = ? AND email = ?`, f.Hash, ctx.Get("email"))
				if err != nil {
					return nil, err
				}
				tmpparts = append(tmpparts, genai.Blob{
					MIMEType: file.MimeType,
					Data:     file.Data,
				})
			}
			cs.History = append(cs.History, &genai.Content{
				Role:  msg.Role,
				Parts: tmpparts,
			})
		} else {
			switch msg.Role {
			case "user":
				cs.History = append(cs.History, &genai.Content{
					Parts: []genai.Part{
						genai.Text(msg.Content),
					},
					Role: "user",
				})
			case "assistant":
				cs.History = append(cs.History, &genai.Content{
					Parts: []genai.Part{
						genai.Text(msg.Content),
					},
					Role: "model",
				})
			}
		}
	}

	return nil, nil
}
