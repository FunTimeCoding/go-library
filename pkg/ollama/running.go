package ollama

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/ollama/ollama/api"
)

func (c *Client) Running() []api.ProcessModelResponse {
	result, e := c.client.ListRunning(c.context)
	errors.PanicOnError(e)

	return result.Models
}
