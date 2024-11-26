package github

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
	"github.com/sashabaranov/go-openai"
	"io"
	"moechat/core/api/part"
	"moechat/core/database"
)

type Client struct {
	resStream *openai.ChatCompletionStream
}

func (c *Client) Ping() {

}

func (c *Client) GetModelList() []string {
	client := resty.New()

	var response []struct {
		Id            string   `json:"id"`
		Name          string   `json:"name"`
		FriendlyName  string   `json:"friendly_name"`
		ModelVersion  int      `json:"model_version"`
		Publisher     string   `json:"publisher"`
		ModelFamily   string   `json:"model_family"`
		ModelRegistry string   `json:"model_registry"`
		License       string   `json:"license"`
		Task          string   `json:"task"`
		Description   string   `json:"description"`
		Summary       string   `json:"summary"`
		Tags          []string `json:"tags"`
	}
	// 发送 GET 请求
	_, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&response).
		Get("https://models.inference.ai.azure.com/models")
	if err != nil {
		return nil
	}
	// 提取 id 字段到 []string 切片
	var ids []string
	for _, model := range response {
		ids = append(ids, model.Name)
	}
	return ids
}

func (c *Client) CreateResStream(ctx echo.Context, completion *part.Completion) error {
	var model database.Model
	if err := database.DB.Get(&model, `SELECT * from model WHERE provider = ? AND active = 1`, completion.Provider); err != nil {
		return err
	}
	config := openai.DefaultConfig(model.APIKey)
	config.BaseURL = model.APIUrl
	client := openai.NewClientWithConfig(config)
	openAIMessages, err := transformToProviderMessages(ctx, completion.Messages)
	if err != nil {
		return err
	}
	request := openai.ChatCompletionRequest{
		Model:    completion.Model,
		Messages: openAIMessages,
		//MaxCompletionTokens: completion.MaxTokens,
		Temperature: completion.Temperature,
		TopP:        completion.TopP,
		Stream:      true,
	}
	//println(completion.MaxTokens)
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

	if len(response.Choices) != 0 {
		n = copy(p, response.Choices[0].Delta.Content)
	}
	return n, nil
}
