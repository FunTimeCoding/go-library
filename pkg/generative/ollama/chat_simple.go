package ollama

import (
	"github.com/funtimecoding/go-library/pkg/generative/ollama/chat_request"
	"github.com/funtimecoding/go-library/pkg/generative/ollama/chat_response"
)

func (c *Client) ChatSimple(p string) *chat_response.Response {
	return c.Chat(chat_request.New().User(p))
}
