package ollama

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/ollama/constant"
	"github.com/funtimecoding/go-library/pkg/ptr"
	"github.com/ollama/ollama/api"
)

func (c *Client) Chat(r *api.ChatRequest) *api.ChatResponse {
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

	return result
}
