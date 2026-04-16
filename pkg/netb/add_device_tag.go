package netb

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) AddDeviceTag(device string, tag string) string {
	result, e := c.client.AddDeviceTag(c.context, device, tag)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
