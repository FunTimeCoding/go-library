package netbox

import (
	"github.com/funtimecoding/go-library/pkg/netbox/device"
	"github.com/funtimecoding/go-library/pkg/netbox/network"
)

func (c *Client) DeviceInterfaceByName(
	d *device.Device,
	name string,
) (*network.Interface, error) {
	result, e := c.DeviceInterfaces(d.Identifier)

	if e != nil {
		return nil, e
	}

	for _, i := range result {
		if i.Name == name {
			return i, nil
		}
	}

	return nil, nil
}
