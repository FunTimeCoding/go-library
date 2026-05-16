package client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) ListTunnelGroups() string {
	result, e := c.client.ListTunnelGroups(c.context)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
