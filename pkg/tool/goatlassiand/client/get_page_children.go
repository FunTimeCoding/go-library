package client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) GetPageChildren(identifier string) string {
	result, e := c.client.GetPageChildren(c.context, identifier)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
