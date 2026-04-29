package netb

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/generated/client"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) CreateInterface(
	device string,
	name string,
	interfaceType string,
) string {
	result, e := c.client.CreateInterface(
		c.context,
		device,
		client.CreateInterfaceRequest{
			Name: name,
			Type: interfaceType,
		},
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
