package ollama

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) Version() string {
	result, e := c.client.Version(c.context)
	errors.PanicOnError(e)

	return result
}
