package part

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
	Files   []struct {
		MimeType string `json:"MimeType"`
		Hash     string `json:"Hash"`
	} `json:"files"`
}

type Completion struct {
	Provider    string    `json:"provider"`
	Model       string    `json:"model"`
	MaxTokens   int       `json:"max_tokens"`
	Temperature float32   `json:"temperature"`
	TopK        int       `json:"topk"`
	TopP        float32   `json:"topp"`
	Messages    []Message `form:"messages" `
}
