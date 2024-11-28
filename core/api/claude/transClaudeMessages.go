package claude

import (
	"encoding/base64"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/liushuangls/go-anthropic/v2"
	"moechat/core/api/part"
	"moechat/core/database"
	"strings"
)

func transformToProviderMessages(ctx echo.Context, msgs []part.Message) ([]anthropic.Message, error) {
	var claudeMessages []anthropic.Message
	var file database.File
	for _, msg := range msgs {
		if msg.Files != nil {
			message := anthropic.Message{
				Role:    anthropic.RoleUser,
				Content: []anthropic.MessageContent{},
			}
			for _, f := range msg.Files {
				if err := database.DB.Get(&file,
					`SELECT * from file WHERE hash = ? AND email = ?`, f.Hash, ctx.Get("email")); err != nil {
					return nil, err
				}
				//todo 没测试过
				if strings.HasPrefix(file.MimeType, "image") {
					message.Content = append(message.Content, anthropic.NewImageMessageContent(
						anthropic.MessageContentSource{
							Type:      anthropic.MessagesContentSourceTypeBase64,
							MediaType: file.MimeType,
							Data:      fmt.Sprintf("data:%s;base64,%s", file.MimeType, base64.StdEncoding.EncodeToString(file.Data)),
						}),
					)
				}

				if !strings.HasPrefix(file.MimeType, "image") {
					message.Content = append(message.Content, anthropic.NewDocumentMessageContent(
						anthropic.MessageContentSource{
							Type:      anthropic.MessagesContentSourceTypeBase64,
							MediaType: file.MimeType,
							Data:      base64.StdEncoding.EncodeToString(file.Data),
						}),
					)
				}

			}
		} else {
			switch msg.Role {
			case "system":
				//claudeMessages = append(claudeMessages, anthropic.NewMultiSystemMessages()...)
			case "user":
				claudeMessages = append(claudeMessages, anthropic.NewUserTextMessage(msg.Content))
			case "assistant":
				claudeMessages = append(claudeMessages, anthropic.NewAssistantTextMessage(msg.Content))
			}
		}
	}

	return claudeMessages, nil
}
