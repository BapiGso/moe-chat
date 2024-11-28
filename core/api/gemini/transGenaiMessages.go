package gemini

import (
	"fmt"
	"github.com/google/generative-ai-go/genai"
	"github.com/labstack/echo/v4"
	"moechat/core/api/part"
	"moechat/core/database"
)

// todo 图片传输
func transformToProviderMessages(ctx echo.Context, msgs []part.Message) ([]*genai.Content, part.Message, error) {
	if len(msgs) == 0 {
		return nil, *new(part.Message), fmt.Errorf("no messages to transform")
	}

	var genaiMessages []*genai.Content

	for _, msg := range msgs { //[:len(msgs)-1] {
		if msg.Files != nil {
			var file database.File
			var tmpparts []genai.Part
			for _, f := range msg.Files {
				err := database.DB.Get(&file,
					`SELECT * from file WHERE hash = ? AND email = ?`, f.Hash, ctx.Get("email"))
				if err != nil {
					return nil, *new(part.Message), err
				}
				tmpparts = append(tmpparts, genai.Blob{
					MIMEType: file.MimeType,
					Data:     file.Data,
				})
			}
			genaiMessages = append(genaiMessages, &genai.Content{
				Role:  msg.Role,
				Parts: tmpparts,
			})
		} else {
			switch msg.Role {
			case "user":
				genaiMessages = append(genaiMessages, &genai.Content{
					Parts: []genai.Part{
						genai.Text(msg.Content),
					},
					Role: "user",
				})
			case "assistant":
				genaiMessages = append(genaiMessages, &genai.Content{
					Parts: []genai.Part{
						genai.Text(msg.Content),
					},
					Role: "model",
				})
			}
		}
	}

	// Process the last message
	lastMsg := msgs[len(msgs)-1]

	return genaiMessages, lastMsg, nil
}
