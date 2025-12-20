package ollama

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/generative/ollama/constant"
	"github.com/funtimecoding/go-library/pkg/generative/ollama/generate_request"
	"github.com/funtimecoding/go-library/pkg/generative/ollama/generate_response"
	"github.com/funtimecoding/go-library/pkg/ptr"
	"github.com/ollama/ollama/api"
)

func (c *Client) Generate(v *generate_request.Request) *generate_response.Response {
	r := v.Get()

	if r.Model == "" {
		r.Model = constant.Llama31
	}

	r.Stream = ptr.To(false)
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

	return generate_response.New(result)
}
