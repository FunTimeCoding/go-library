package client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) ListClusterTypes() string {
	result, e := c.client.ListClusterTypes(c.context)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
