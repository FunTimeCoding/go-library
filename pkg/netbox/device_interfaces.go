package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/network"
)

func (c *Client) DeviceInterfaces(device int32) []*network.Interface {
	result, _, e := c.client.DcimAPI.DcimInterfacesList(
		c.context,
	).DeviceId([]int32{device}).Execute()
	errors.PanicOnError(e)

	return network.NewSlice(result.Results)
}
