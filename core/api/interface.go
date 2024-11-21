package api

import (
	"moechat/core/api/azure"
	"moechat/core/api/github"
	"moechat/core/api/openai"
)

type Adapter interface {
	GetModelList() []string
	Ping()
}

func New(AiType string) Adapter {
	switch AiType {
	case "azure":
		return Adapter(new(azure.Azure))
	case "GitHub":
		return Adapter(new(github.GitHub))
	case "OpenAI":
		return Adapter(new(openai.Openai))
	default:
		return nil
	}
}
