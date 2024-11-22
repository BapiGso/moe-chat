package api

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"moechat/core/api/github"
)

type Adapter interface {
	GetModelList() []string
	Ping()
	ReadResSteam(p []byte) (n int, err error)
	CreateResSteam(ctx echo.Context, baseModel string, message json.RawMessage) error
	//FromProviderFormat(providerMsg interface{}) ([]UnifiedMessage, error)
}

func New(provider string) Adapter {
	switch provider {
	case "Azure":
	//return new(azure.Client)
	case "Claude":
		//return new(claude.Client)
	case "GitHub":
		return new(github.Client)
	case "OpenAI":
		//return new(openai.Openai)
	}
	return nil
}

// UnifiedMessage 统一的消息格式，用于内部处理
