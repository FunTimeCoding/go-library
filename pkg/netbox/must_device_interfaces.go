package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/network"
)

func (c *Client) MustDeviceInterfaces(device int32) []*network.Interface {
	result, e := c.DeviceInterfaces(device)
	errors.PanicOnError(e)

	return result
}
