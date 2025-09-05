package netbox

import (
	"github.com/funtimecoding/go-library/pkg/netbox/device"
	"github.com/funtimecoding/go-library/pkg/netbox/network"
	"log"
)

func (c *Client) DeviceInterfaceByNameStrict(
	d *device.Device,
	name string,
) *network.Interface {
	for _, i := range c.DeviceInterfaces(d.Identifier) {
		if i.Name == name {
			return i
		}
	}

	log.Panicf(
		"interface %s not found for device %s",
		name,
		d.Name,
	)

	return network.Stub()
}
