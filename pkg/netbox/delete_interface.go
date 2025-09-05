package netbox

import "github.com/funtimecoding/go-library/pkg/netbox/device"

func (c *Client) DeleteInterface(
	d *device.Device,
	name string,
) {
	r, e := c.client.DcimAPI.DcimInterfacesDestroy(
		c.context,
		c.DeviceInterfaceByName(d, name, true).Identifier,
	).Execute()
	verifyDelete("interface", r, e)
}
