package netb

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) ListInterfaces(name string) string {
	result, e := c.client.ListInterfaces(c.context, name)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
