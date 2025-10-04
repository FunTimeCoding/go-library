package openai

import (
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func New(token string) *Client {
	return &Client{client: openai.NewClient(option.WithAPIKey(token))}
}
