package ollama

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/generative/ollama/generate_request"
	"github.com/funtimecoding/go-library/pkg/generative/ollama/generate_response"
)

func (c *Client) MustGenerate(v *generate_request.Request) *generate_response.Response {
	result, e := c.Generate(v)
	errors.PanicOnError(e)

	return result
}
