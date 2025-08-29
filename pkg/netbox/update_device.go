package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/device"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) updateDevice(
	d *device.Device,
	r *netbox.PatchedWritableDeviceWithConfigContextRequest,
) *device.Device {
	result, _, e := c.client.DcimAPI.DcimDevicesPartialUpdate(
		c.context,
		d.Identifier,
	).PatchedWritableDeviceWithConfigContextRequest(*r).Execute()
	errors.PanicOnError(e)

	return device.New(result)
}
