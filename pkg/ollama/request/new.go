package request

import "github.com/ollama/ollama/api"

func New(prompt string) *api.GenerateRequest {
	return &api.GenerateRequest{Prompt: prompt}
}
