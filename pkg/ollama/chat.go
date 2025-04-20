package ollama

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/ollama/chat_request"
	"github.com/funtimecoding/go-library/pkg/ollama/chat_response"
	"github.com/funtimecoding/go-library/pkg/ollama/constant"
	"github.com/funtimecoding/go-library/pkg/ptr"
	"github.com/ollama/ollama/api"
)

func (c *Client) Chat(v *chat_request.Request) *chat_response.Response {
	r := v.Get()

	if r.Model == "" {
		r.Model = constant.Llama31
	}

	r.Stream = ptr.To[bool](false)
	var result *api.ChatResponse
	errors.PanicOnError(
		c.client.Chat(
			c.context,
			r,
			func(r api.ChatResponse) error {
				result = &r

				return nil
			},
		),
	)

	return chat_response.New(result)
}
