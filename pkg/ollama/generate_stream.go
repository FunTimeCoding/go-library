package ollama

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/ollama/constant"
	"github.com/funtimecoding/go-library/pkg/ptr"
	"github.com/ollama/ollama/api"
)

func (c *Client) GenerateStream(
	r *api.GenerateRequest,
	output chan<- string,
	response *api.GenerateResponse,
) {
	if r.Model == "" {
		r.Model = constant.Llama31
	}

	r.Stream = ptr.To[bool](true)
	errors.PanicOnError(
		c.client.Generate(
			c.context,
			r,
			func(r api.GenerateResponse) error {
				if r.Done {
					response.Response = r.Response
					response.Metrics = r.Metrics
					response.Done = r.Done
					response.Model = r.Model
					response.CreatedAt = r.CreatedAt
				}

				output <- r.Response

				return nil
			},
		),
	)

	close(output)
}
