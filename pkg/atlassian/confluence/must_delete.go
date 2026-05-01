package confluence

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) MustDelete(pageIdentifier string) {
	errors.PanicOnError(c.Delete(pageIdentifier))
}
