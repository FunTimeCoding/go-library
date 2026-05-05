package client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) ListSites() string {
	result, e := c.client.ListSites(c.context)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
