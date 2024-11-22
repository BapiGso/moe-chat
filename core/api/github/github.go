package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sashabaranov/go-openai"
	"io"
	"moechat/core/database"
)

type Client struct {
	resStream *openai.ChatCompletionStream
}

func (g *Client) Ping() {

}

func (g *Client) GetModelList() []string {

	return []string{"gpt-4o-mini", "gpt-4o"}
}

func (g *Client) CreateResSteam(ctx echo.Context, baseModel string, msgs json.RawMessage) error {
	openAIMessages, err := transformToProviderMessages(msgs)
	if err != nil {
		return err
	}
	request := openai.ChatCompletionRequest{
		Model:    baseModel,
		Messages: openAIMessages,
		Stream:   true,
	}
	var model database.Model

	if err := database.DB.Get(&model, `SELECT * from model WHERE provider = 'GitHub'`); err != nil {
		return err
	}
	client := openai.NewClientWithConfig(openai.DefaultAzureConfig(model.APIKey, model.APIUrl))
	g.resStream, err = client.CreateChatCompletionStream(ctx.Request().Context(), request)
	return err
}

func (g *Client) ReadResSteam(p []byte) (n int, err error) {
	// receive new response from the stream
	response, err := g.resStream.Recv()
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
