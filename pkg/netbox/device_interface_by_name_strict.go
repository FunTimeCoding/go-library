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
	for _, element := range c.DeviceInterfaces(d.Identifier) {
		if element.Name == name {
			return element
		}
	}

	log.Panicf(
		"interface %s not found for device %s",
		name,
		d.Name,
	)

	return network.Stub()
}
