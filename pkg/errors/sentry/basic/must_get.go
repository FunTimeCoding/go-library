package basic

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) MustGet(
	path string,
	query map[string]string,
) []byte {
	result, e := c.Get(path, query)
	errors.PanicOnError(e)

	return result
}
