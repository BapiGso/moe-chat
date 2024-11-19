package openai

import (
	"context"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
)

func init() {
	example()
}

func example() {
	config := openai.DefaultAzureConfig("sk-esGX9mNr6eW5ptPfLzjMz8DfwgNVrPqwlY6vg2sC9KtSh9E3", "https://api.deepseek.com")
	// If you use a deployment name different from the model name, you can customize the AzureModelMapperFunc function
	//config.AzureModelMapperFunc = func(model string) string {
	//	azureModelMapping := map[string]string{
	//		"gpt-4o": "gpt-4o",
	//	}
	//	return azureModelMapping[model]
	//}

	client := openai.NewClientWithConfig(config)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: "deepseek-chat",
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Hello Azure OpenAI!",
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}

func example2() {

	client := openai.NewClient("sk-Ho9hOFUHszHcItvsb1200AHQExqAqnvOuJXNW9xHGLmSddLn")
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4oMini,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Hello!",
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}
