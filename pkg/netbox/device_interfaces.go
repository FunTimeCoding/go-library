package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/network"
)

func (c *Client) DeviceInterfaces(device string) []*network.Interface {
	result, _, e := c.client.DcimAPI.DcimInterfacesList(
		c.context,
	).Device([]*string{&device}).Execute()
	errors.PanicOnError(e)

	return network.NewSlice(result.Results, c.interfaceTypes)
}
