package client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) ListInstances() string {
	result, e := c.client.ListInstancesWithResponse(c.context)
	errors.PanicOnError(e)

	return web.ReadString(result.HTTPResponse)
}
