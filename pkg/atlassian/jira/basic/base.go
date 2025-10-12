package basic

import "github.com/funtimecoding/go-library/pkg/web/locator"

func (c *Client) Base() *locator.Locator {
	return c.base
}
