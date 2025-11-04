package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/device"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) updateDevice(
	d *device.Device,
	q *netbox.PatchedWritableDeviceWithConfigContextRequest,
) *device.Device {
	result, r, e := c.client.DcimAPI.DcimDevicesPartialUpdate(
		c.context,
		d.Identifier,
	).PatchedWritableDeviceWithConfigContextRequest(*q).Execute()
	errors.PanicOnWebError(r, e)

	return device.New(result)
}
