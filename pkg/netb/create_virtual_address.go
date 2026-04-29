package netb

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/generated/client"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) CreateVirtualAddress(
	vmName string,
	interfaceName string,
	address string,
) string {
	result, e := c.client.CreateVirtualAddress(
		c.context,
		vmName,
		client.CreateAddressRequest{
			Interface: interfaceName,
			Address:   address,
		},
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
