package ollama

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/ollama/ollama/api"
)

func (c *Client) MustShow(model string) *api.ShowResponse {
	result, e := c.Show(model)
	errors.PanicOnError(e)

	return result
}
