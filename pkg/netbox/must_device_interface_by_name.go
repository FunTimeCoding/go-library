package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/device"
	"github.com/funtimecoding/go-library/pkg/netbox/network"
)

func (c *Client) MustDeviceInterfaceByName(
	d *device.Device,
	name string,
) *network.Interface {
	result, e := c.DeviceInterfaceByName(d, name)
	errors.PanicOnError(e)

	return result
}
