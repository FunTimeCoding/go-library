package netbox

import "github.com/funtimecoding/go-library/pkg/netbox/device"

func (c *Client) DeleteInterfaceByName(
	d *device.Device,
	name string,
) {
	c.DeleteInterface(c.DeviceInterfaceByNameStrict(d, name).Identifier)
}
