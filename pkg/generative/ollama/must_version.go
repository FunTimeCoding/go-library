package ollama

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) MustVersion() string {
	result, e := c.Version()
	errors.PanicOnError(e)

	return result
}
