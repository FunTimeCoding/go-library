package netb

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/client"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) CreateVirtualInterface(
	vmName string,
	name string,
) string {
	result, e := c.client.CreateVirtualInterface(
		c.context,
		vmName,
		client.CreateVirtualInterfaceRequest{Name: name},
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
