package proxmox

import "github.com/funtimecoding/go-library/pkg/errors"

func MustGet[T any](
	c *Client,
	locator string,
	result *T,
) {
	errors.PanicOnError(Get(c, locator, result))
}
