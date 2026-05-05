package client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) CreateDeviceRole(name string) string {
	result, e := c.client.CreateDeviceRole(
		c.context,
		client.CreateNameRequest{Name: name},
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
