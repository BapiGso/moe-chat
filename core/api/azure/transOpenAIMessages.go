package azure

import (
	"encoding/base64"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sashabaranov/go-openai"
	"moechat/core/api/part"
	"moechat/core/database"
	"strings"
)

func transformToProviderMessages(ctx echo.Context, msgs []part.Message) ([]openai.ChatCompletionMessage, error) {
	var openAIMessages []openai.ChatCompletionMessage
	var file database.File
	for _, msg := range msgs {
		if msg.Files != nil {
			var MultiContent []openai.ChatMessagePart
			for _, f := range msg.Files {
				if err := database.DB.Get(&file,
					`SELECT * from file WHERE hash = ? AND email = ?`, f.Hash, ctx.Get("email")); err != nil {
					return nil, err
				}
				if !strings.HasPrefix(file.MimeType, "image") {
					continue
				}
				MultiContent = append(MultiContent, openai.ChatMessagePart{
					Type: openai.ChatMessagePartTypeImageURL,
					ImageURL: &openai.ChatMessageImageURL{
						URL: fmt.Sprintf("data:%s;base64,%s", file.MimeType, base64.StdEncoding.EncodeToString(file.Data)),
						//URL: string(file.Data),
					},
				})
			}
			openAIMessages = append(openAIMessages, openai.ChatCompletionMessage{
				Role:         msg.Role,
				MultiContent: MultiContent,
			})
		} else {
			openAIMessages = append(openAIMessages, openai.ChatCompletionMessage{
				Role:    msg.Role,
				Content: msg.Content,
			})
		}

	}
	return openAIMessages, nil
}
