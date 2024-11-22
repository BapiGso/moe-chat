package github

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"io"
)

func CreateOpenAIResponseReader(ctx context.Context) (io.Reader, error) {
	request := openai.ChatCompletionRequest{
		Model: openai.GPT4oMini,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "帮我写一个500字的抒情文",
			},
		},
		Stream: true,
	}
	client := openai.NewClientWithConfig(openai.DefaultAzureConfig("", "https://models.inference.ai.azure.com"))
	stream, err := client.CreateChatCompletionStream(ctx, request)
	if err != nil {
		return nil, err
	}

	return &GitHub{stream: stream}, nil
}
