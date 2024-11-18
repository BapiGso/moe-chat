package claude

import (
	"context"
	"errors"
	"fmt"
	"github.com/liushuangls/go-anthropic/v2"
)

func example() {
	client := anthropic.NewClient("your anthropic api key")
	resp, err := client.CreateMessages(context.Background(), anthropic.MessagesRequest{
		Model: anthropic.ModelClaude3Haiku20240307,
		Messages: []anthropic.Message{
			anthropic.NewUserTextMessage("What is your name?"),
		},
		MaxTokens: 1000,
	})
	if err != nil {
		var e *anthropic.APIError
		if errors.As(err, &e) {
			fmt.Printf("Messages error, type: %s, message: %s", e.Type, e.Message)
		} else {
			fmt.Printf("Messages error: %v\n", err)
		}
		return
	}
	fmt.Println(resp.Content[0].GetText())
}
