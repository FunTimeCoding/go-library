package openai

import (
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

func New(token string) *Client {
	return &Client{client: openai.NewClient(option.WithAPIKey(token))}
}
