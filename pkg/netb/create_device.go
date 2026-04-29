package netb

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/generated/client"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) CreateDevice(
	name string,
	role string,
	deviceType string,
	site string,
	tenant *string,
) string {
	result, e := c.client.CreateDevice(
		c.context,
		client.CreateDeviceRequest{
			Name:   name,
			Role:   role,
			Type:   deviceType,
			Site:   site,
			Tenant: tenant,
		},
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
