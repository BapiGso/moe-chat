package grok

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
	"github.com/sashabaranov/go-openai"
	"io"
	"moechat/core/database"
)

type Client struct {
	resStream *openai.ChatCompletionStream
}

func (c *Client) Ping() {

}

func (c *Client) GetModelList() []string {
	var model database.Model
	if err := database.DB.Get(&model, `SELECT * from model WHERE provider = 'Grok' AND active = 1`); err != nil {
		return nil
	}
	client := resty.New()

	var response struct {
		Data []struct {
			ID      string `json:"id"`
			Created int64  `json:"created"`
			Object  string `json:"object"`
			OwnedBy string `json:"owned_by"`
		} `json:"data"`
		Object string `json:"object"`
	}
	// 发送 GET 请求
	_, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+model.APIKey).
		SetResult(&response).
		Get("https://api.x.ai/v1/models")
	if err != nil {
		return nil
	}
	// 提取 id 字段到 []string 切片
	var ids []string
	for _, model := range response.Data {
		ids = append(ids, model.ID)
	}
	return ids
}

func (c *Client) CreateResStream(ctx echo.Context, baseModel string, msgs json.RawMessage) error {
	var model database.Model
	if err := database.DB.Get(&model, `SELECT * from model WHERE provider = 'Grok'`); err != nil {
		return err
	}
	config := openai.DefaultConfig(model.APIKey)
	config.BaseURL = model.APIUrl
	client := openai.NewClientWithConfig(config)
	openAIMessages, err := transformToProviderMessages(msgs)
	if err != nil {
		return err
	}
	request := openai.ChatCompletionRequest{
		Model:    baseModel,
		Messages: openAIMessages,
		Stream:   true,
	}
	c.resStream, err = client.CreateChatCompletionStream(ctx.Request().Context(), request)
	return err
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

	// Copy the content directly to the provided buffer
	n = copy(p, response.Choices[0].Delta.Content)
	return n, nil
}
