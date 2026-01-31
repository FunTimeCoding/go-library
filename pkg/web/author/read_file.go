package author

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) ReadFile(path string) []byte {
	result, e := c.client.Read(path)
	errors.PanicOnError(e)

	return result
}
