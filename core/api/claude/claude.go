package claude

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/liushuangls/go-anthropic/v2"
	"moechat/core/api/part"
	"moechat/core/database"
)

type Client struct {
	resStream     anthropic.MessagesResponse
	MessagesEvent chan *string
}

func (c *Client) GetModelList() []string {
	return nil
}

func (c *Client) Read(p []byte) (n int, err error) {
	select {
	case response := <-c.MessagesEvent:
		n = copy(p, *response)
	}
	//response := c.resStream.Content
	return n, nil
}

func (c *Client) CreateResStream(ctx echo.Context, completion *part.Completion) error {
	var model database.Model
	var err error
	if err := database.DB.Get(&model, `SELECT * from model WHERE provider = 'Claude' AND active = 1`); err != nil {
		return err
	}
	client := anthropic.NewClient(model.APIKey)
	claudeMessages, err := transformToProviderMessages(ctx, completion.Messages)
	if err != nil {
		return err
	}
	request := anthropic.MessagesRequest{
		Model:       anthropic.Model(completion.Model),
		Messages:    claudeMessages,
		MaxTokens:   completion.MaxTokens,
		Stream:      true,
		Temperature: &completion.Temperature,
		TopP:        &completion.TopP,
		TopK:        &completion.TopK,
	}
	c.resStream, err = client.CreateMessagesStream(ctx.Request().Context(), anthropic.MessagesStreamRequest{
		MessagesRequest: request,
		OnContentBlockDelta: func(data anthropic.MessagesEventContentBlockDeltaData) {
			c.MessagesEvent <- data.Delta.Text
			//fmt.Printf("Stream Content: %s\n", data.Delta.Text)
		},
	})
	if err != nil {
		var e *anthropic.APIError
		if errors.As(err, &e) {
			fmt.Printf("Messages stream error, type: %s, message: %s", e.Type, e.Message)
		} else {
			fmt.Printf("Messages stream error: %v\n", err)
		}
		return e
	}
	return err

}

func (c *Client) Ping() {

}
