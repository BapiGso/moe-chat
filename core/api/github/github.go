package github

import (
	"context"
	"errors"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"io"
)

type GitHub struct {
	stream *openai.ChatCompletionStream
}

func (g *GitHub) Ping() {

}

func (g *GitHub) GetModelList() []string {

	return []string{"gpt-4o-mini", "gpt-4o"}
}

func (r *GitHub) Read(p []byte) (n int, err error) {
	// receive new response from the stream
	response, err := r.stream.Recv()
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

func (g *GitHub) Test() *openai.ChatCompletionStream {
	ctx := context.Background()
	req := openai.ChatCompletionRequest{
		Model: openai.GPT4o,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "Hello Azure OpenAI!",
			},
		},
		Stream: true,
	}
	client := openai.NewClientWithConfig(openai.DefaultAzureConfig("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", "https://models.inference.ai.azure.com"))
	stream, err := client.CreateChatCompletionStream(ctx, req)

	if err != nil {
	}
	defer stream.Close()

	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("\nStream finished")
		}
		if err != nil {
			fmt.Printf("\nStream error: %v\n", err)

		}
		fmt.Printf(response.Choices[0].Delta.Content)
	}
	return nil
}
