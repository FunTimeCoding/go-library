package netbox

import (
	"github.com/funtimecoding/go-library/pkg/netbox/device"
	"github.com/funtimecoding/go-library/pkg/netbox/network"
	"log"
)

func (c *Client) DeviceInterfaceByName(
	d *device.Device,
	name string,
	panic bool,
) *network.Interface {
	for _, element := range c.DeviceInterfaces(d.Name) {
		if element.Name == name {
			return element
		}
	}

	if panic {
		log.Panicf(
			"interface %s not found for device %s",
			name,
			d.Name,
		)
	}

	return nil
}
