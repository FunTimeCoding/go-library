package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/device"
)

func (c *Client) MustDevicesByMatch(s string) []*device.Device {
	result, e := c.DevicesByMatch(s)
	errors.PanicOnError(e)

	return result
}
