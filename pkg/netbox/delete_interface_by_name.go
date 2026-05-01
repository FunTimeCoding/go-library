package netbox

import "github.com/funtimecoding/go-library/pkg/netbox/device"

func (c *Client) DeleteInterfaceByName(
	d *device.Device,
	name string,
) error {
	i, e := c.DeviceInterfaceByNameStrict(d, name)

	if e != nil {
		return e
	}

	return c.DeleteInterface(i.Identifier)
}
