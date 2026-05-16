package client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) ListTunnels() string {
	result, e := c.client.ListTunnels(c.context)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
