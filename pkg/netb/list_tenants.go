package netb

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) ListTenants() string {
	result, e := c.client.ListTenants(c.context)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
