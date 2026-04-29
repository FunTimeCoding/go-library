package ollama

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/generative/ollama/chat_request"
	"github.com/funtimecoding/go-library/pkg/generative/ollama/chat_response"
)

func (c *Client) MustChat(v *chat_request.Request) *chat_response.Response {
	result, e := c.Chat(v)
	errors.PanicOnError(e)

	return result
}
