package api

import (
	"errors"
	"github.com/labstack/echo/v4"
	"moechat/core/api/azure"
	"moechat/core/api/claude"
	"moechat/core/api/gemini"
	"moechat/core/api/github"
	"moechat/core/api/grok"
	"moechat/core/api/ollama"
	"moechat/core/api/openai"
	"moechat/core/api/part"
)

type Adapter interface {
	GetModelList() []string
	CreateResStream(ctx echo.Context, completion *part.Completion) error
	Read(p []byte) (n int, err error)
}

func New(provider string) (Adapter, error) {
	switch provider {
	case "Azure":
		return new(azure.Client), nil
	case "Claude":
		return new(claude.Client), nil
	case "Gemini":
		return new(gemini.Client), nil
	case "GitHub":
		return new(github.Client), nil
	case "Grok":
		return new(grok.Client), nil
	case "Ollama":
		return new(ollama.Client), nil
	case "OpenAI":
		return new(openai.Client), nil
	default:
		return nil, errors.New("unsupported provider")
	}
}
