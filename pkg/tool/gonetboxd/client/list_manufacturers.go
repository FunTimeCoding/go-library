package client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) ListManufacturers() string {
	result, e := c.client.ListManufacturers(c.context)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
