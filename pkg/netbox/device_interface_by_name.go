package netbox

import (
	"github.com/funtimecoding/go-library/pkg/netbox/device"
	"github.com/funtimecoding/go-library/pkg/netbox/network"
)

func (c *Client) DeviceInterfaceByName(
	d *device.Device,
	name string,
) *network.Interface {
	for _, i := range c.DeviceInterfaces(d.Identifier) {
		if i.Name == name {
			return i
		}
	}

	return nil
}
