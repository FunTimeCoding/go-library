package client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) GetNodeStatus(name string) string {
	result, e := c.client.GetNodeStatus(c.context, name)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
