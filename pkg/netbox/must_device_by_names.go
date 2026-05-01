package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/device"
)

func (c *Client) MustDeviceByNames(n []string) *device.Device {
	result, e := c.DeviceByNames(n)
	errors.PanicOnError(e)

	return result
}
