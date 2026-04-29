package basic

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) MustDelete(path string) {
	errors.PanicOnError(c.Delete(path))
}
