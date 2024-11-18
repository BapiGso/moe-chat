package openai

import (
	"context"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
)

func init() {
	//example()
}

func example() {
	config := openai.DefaultAzureConfig("", "https://models.inference.ai.azure.com")
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
			Model: openai.GPT4o,
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
