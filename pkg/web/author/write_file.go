package author

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) WriteFile(
	path string,
	content []byte,
) {
	errors.PanicOnError(c.client.Write(path, content, 0644))
}
