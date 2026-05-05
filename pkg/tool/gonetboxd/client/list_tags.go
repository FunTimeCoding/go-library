package client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) ListTags() string {
	result, e := c.client.ListTags(c.context)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
