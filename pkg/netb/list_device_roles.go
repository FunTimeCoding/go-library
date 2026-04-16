package netb

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) ListDeviceRoles() string {
	result, e := c.client.ListDeviceRoles(c.context)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
