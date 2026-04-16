package netb

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/client"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) CreateAddress(
	device string,
	interfaceName string,
	address string,
) string {
	result, e := c.client.CreateAddress(
		c.context,
		device,
		client.CreateAddressRequest{
			Interface: interfaceName,
			Address:   address,
		},
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
