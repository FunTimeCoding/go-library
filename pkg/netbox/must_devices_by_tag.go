package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/device"
)

func (c *Client) MustDevicesByTag(s string) []*device.Device {
	result, e := c.DevicesByTag(s)
	errors.PanicOnError(e)

	return result
}
