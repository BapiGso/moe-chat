package ollama

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
	"github.com/sashabaranov/go-openai"
	"io"
	"moechat/core/database"
	"net/url"
)

type Client struct {
	resStream *openai.ChatCompletionStream
}

func (c *Client) Ping() {

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

func (c *Client) GetModelList() []string {
	var model database.Model
	if err := database.DB.Get(&model, `SELECT * from model WHERE provider = 'Ollama'`); err != nil {
		return nil
	}
	parsedURL, err := url.Parse(model.APIUrl)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return nil
	}

	// 提取 Scheme 和 Host
	baseURL := parsedURL.Scheme + "://" + parsedURL.Host
	client := resty.New()
	var response struct {
		Object string `json:"object"`
		Data   []struct {
			Id      string `json:"id"`
			Object  string `json:"object"`
			Created int    `json:"created"`
			OwnedBy string `json:"owned_by"`
		} `json:"data"`
	}
	// 发送 GET 请求
	if _, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&response).
		Get(baseURL + "/v1/models"); err != nil {
		return nil
	}
	// 提取 id 字段到 []string 切片
	var ids []string
	for _, model := range response.Data {
		ids = append(ids, model.Id)
	}
	return ids
}

func (c *Client) CreateResStream(ctx echo.Context, baseModel string, msgs json.RawMessage) error {
	var model database.Model
	if err := database.DB.Get(&model, `SELECT * from model WHERE provider = 'Ollama' AND active = 1`); err != nil {
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
