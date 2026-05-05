package client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) CreateTenant(name string) string {
	result, e := c.client.CreateTenant(
		c.context,
		client.CreateNameRequest{Name: name},
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
