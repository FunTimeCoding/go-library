package ollama

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/ollama/ollama/api"
	"k8s.io/utils/pointer"
)

func (c *Client) Generate(r *api.GenerateRequest) *api.GenerateResponse {
	if r.Model == "" {
		r.Model = "llama3.1"
	}

	if r.Stream == nil {
		r.Stream = pointer.Bool(false)
	}

	var result *api.GenerateResponse
	errors.PanicOnError(
		c.client.Generate(
			context.Background(),
			r,
			func(r api.GenerateResponse) error {
				result = &r

				return nil
			},
		),
	)

	return result
}
