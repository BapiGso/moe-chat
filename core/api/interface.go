package api

import "moechat/core/api/azure"

type Adapter interface {
	GetModelList()
	Ping()
}

func New(AiType string) Adapter {
	switch AiType {
	case "azure":
		return Adapter(new(azure.Azure))
	}
	return nil
}
