package ollama

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/ollama/ollama/api"
)

func (c *Client) List() []api.ListModelResponse {
	result, e := c.client.List(c.context)
	errors.PanicOnError(e)

	return result.Models
}
