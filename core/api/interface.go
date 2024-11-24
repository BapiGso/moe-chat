package api

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"moechat/core/api/azure"
	"moechat/core/api/claude"
	"moechat/core/api/gemini"
	"moechat/core/api/github"
	"moechat/core/api/grok"
	"moechat/core/api/ollama"
	"moechat/core/api/openai"
)

type Adapter interface {
	GetModelList() []string
	CreateResStream(ctx echo.Context, baseModel string, msgs json.RawMessage) error
	Read(p []byte) (n int, err error)
}

func New(provider string) Adapter {
	switch provider {
	case "Azure":
		return new(azure.Client)
	case "Claude":
		return new(claude.Client)
	case "Gemini":
		return new(gemini.Client)
	case "GitHub":
		return new(github.Client)
	case "Grok":
		return new(grok.Client)
	case "Ollama":
		return new(ollama.Client)
	case "OpenAI":
		return new(openai.Client)
	}
	return nil
}
