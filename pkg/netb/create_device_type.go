package netb

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/client"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) CreateDeviceType(model string, manufacturer string) string {
	result, e := c.client.CreateDeviceType(
		c.context,
		client.CreateDeviceTypeRequest{
			Model:        model,
			Manufacturer: manufacturer,
		},
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
