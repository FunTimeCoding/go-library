package ollama

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/ollama/ollama/api"
)

func (c *Client) List() []api.ListModelResponse {
	r, e := c.client.List(c.context)
	errors.PanicOnError(e)

	return r.Models
}
