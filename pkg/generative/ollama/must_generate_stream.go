package ollama

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/ollama/ollama/api"
)

func (c *Client) MustGenerateStream(
	r *api.GenerateRequest,
	output chan<- string,
	response *api.GenerateResponse,
) {
	errors.PanicOnError(c.GenerateStream(r, output, response))
}
