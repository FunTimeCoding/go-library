package netbox

import "github.com/funtimecoding/go-library/pkg/netbox/device"

func (c *Client) DeleteInterface(
	d *device.Device,
	name string,
) {
	r, e := c.client.DcimAPI.DcimInterfacesDestroy(
		c.context,
		c.DeviceInterfaceByNameStrict(d, name).Identifier,
	).Execute()
	verifyDelete("interface", r, e)
}
