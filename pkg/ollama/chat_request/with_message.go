package chat_request

import (
	"github.com/funtimecoding/go-library/pkg/ollama/constant"
	"github.com/ollama/ollama/api"
)

func WithMessage(content string) *api.ChatRequest {
	return &api.ChatRequest{
		Messages: []api.Message{
			{
				Role:    constant.UserRole,
				Content: content,
			},
		},
	}
}
