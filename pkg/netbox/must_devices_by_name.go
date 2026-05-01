package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/device"
)

func (c *Client) MustDevicesByName(s string) []*device.Device {
	result, e := c.DevicesByName(s)
	errors.PanicOnError(e)

	return result
}
