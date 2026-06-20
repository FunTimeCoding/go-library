package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/device"
)

func (c *Client) MustDeviceInterfaceExists(
	d *device.Device,
	name string,
) bool {
	result, e := c.DeviceInterfaceExists(d, name)
	errors.PanicOnError(e)

	return result
}
