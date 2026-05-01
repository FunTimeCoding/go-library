package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/device"
)

func (c *Client) MustDeviceByName(n string) *device.Device {
	result, e := c.DeviceByName(n)
	errors.PanicOnError(e)

	return result
}
