package proxmox

import "github.com/funtimecoding/go-library/pkg/errors"

func Get[T any](
	c *Client,
	locator string,
	result *T,
) {
	errors.PanicOnError(c.client.Get(c.context, locator, result))
}
