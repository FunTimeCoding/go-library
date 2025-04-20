package ollama

import (
	"github.com/funtimecoding/go-library/pkg/ollama/chat_request"
	"github.com/funtimecoding/go-library/pkg/ollama/chat_response"
)

func (c *Client) ChatSimple(prompt string) *chat_response.Response {
	return c.Chat(chat_request.New().User(prompt).Get())
}
