package ollama

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/ollama/constant"
	"github.com/funtimecoding/go-library/pkg/ptr"
	"github.com/ollama/ollama/api"
)

func (c *Client) Generate(r *api.GenerateRequest) *api.GenerateResponse {
	if r.Model == "" {
		r.Model = constant.Llama31
	}

	if r.Stream == nil {
		r.Stream = ptr.To[bool](false)
	}

	var result *api.GenerateResponse
	errors.PanicOnError(
		c.client.Generate(
			c.context,
			r,
			func(r api.GenerateResponse) error {
				result = &r

				return nil
			},
		),
	)

	return result
}
