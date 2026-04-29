package basic

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) MustPut(
	path string,
	body any,
) []byte {
	result, e := c.Put(path, body)
	errors.PanicOnError(e)

	return result
}
