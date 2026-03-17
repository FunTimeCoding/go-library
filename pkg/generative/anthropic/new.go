package anthropic

import (
	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
)

func New(token string) *Client {
	return &Client{client: anthropic.NewClient(option.WithAPIKey(token))}
}
