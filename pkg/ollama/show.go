package ollama

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/ollama/ollama/api"
)

func (c *Client) Show(model string) *api.ShowResponse {
	result, e := c.client.Show(c.context, &api.ShowRequest{Model: model})
	errors.PanicOnError(e)

	return result
}
