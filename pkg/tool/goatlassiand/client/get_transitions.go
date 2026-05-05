package client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) GetTransitions(key string) string {
	result, e := c.client.GetTransitions(c.context, key)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
