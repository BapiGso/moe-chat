package openai

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sashabaranov/go-openai"
	"io"
	"moechat/core/api/part"
	"moechat/core/database"
)

type Client struct {
	resStream *openai.ChatCompletionStream
}

func (c *Client) Read(p []byte) (n int, err error) {
	// receive new response from the stream
	response, err := c.resStream.Recv()
	if err != nil {
		if errors.Is(err, io.EOF) {
			return 0, io.EOF
		}
		return 0, fmt.Errorf("stream error: %w", err)
	}

	if len(response.Choices) != 0 {
		n = copy(p, response.Choices[0].Delta.Content)
	}
	return n, nil
}

func (c *Client) CreateResStream(ctx echo.Context, completion *part.Completion) error {
	var err error
	var model database.Model
	if err := database.DB.Get(&model, `SELECT * from model WHERE provider = ? AND active = 1`, completion.Provider); err != nil {
		return err
	}
	client := openai.NewClient(model.APIKey)
	openAIMessages, err := transformToProviderMessages(ctx, completion.Messages)
	request := openai.ChatCompletionRequest{
		Model:       completion.Model,
		Messages:    openAIMessages,
		MaxTokens:   completion.MaxTokens,
		Temperature: completion.Temperature,
		TopP:        completion.TopP,
		Stream:      true,
	}
	c.resStream, err = client.CreateChatCompletionStream(ctx.Request().Context(), request)
	return err
}

func (c *Client) Ping() {

}

func (c *Client) GetModelList() []string {
	return nil
}
