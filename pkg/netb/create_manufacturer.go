package netb

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/generated/client"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) CreateManufacturer(name string) string {
	result, e := c.client.CreateManufacturer(
		c.context,
		client.CreateNameRequest{Name: name},
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
