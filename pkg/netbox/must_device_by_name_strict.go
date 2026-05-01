package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/device"
)

func (c *Client) MustDeviceByNameStrict(n string) *device.Device {
	result, e := c.DeviceByNameStrict(n)
	errors.PanicOnError(e)

	return result
}
