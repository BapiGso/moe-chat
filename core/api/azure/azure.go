package azure

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
	response, err := c.resStream.Recv()
	if err != nil {
		if errors.Is(err, io.EOF) {
			return 0, io.EOF
		}
		return 0, fmt.Errorf("stream error: %w", err)
	}

	// Copy the content directly to the provided buffer
	n = copy(p, response.Choices[0].Delta.Content)
	return n, nil
}

func (c *Client) CreateResStream(ctx echo.Context, baseModel string, msgs []part.Message) error {
	var err error
	var model database.Model
	if err := database.DB.Get(&model, `SELECT * from model WHERE provider = 'Azure' AND active = 1`); err != nil {
		return err
	}
	config := openai.DefaultAzureConfig(model.APIKey, model.APIUrl)
	client := openai.NewClientWithConfig(config)
	openAIMessages, err := transformToProviderMessages(msgs)
	request := openai.ChatCompletionRequest{
		Model:    baseModel,
		Messages: openAIMessages,
		Stream:   true,
	}
	c.resStream, err = client.CreateChatCompletionStream(ctx.Request().Context(), request)
	return err
}

func (c *Client) Ping() {

}

func (c *Client) GetModelList() []string {
	return nil
}
