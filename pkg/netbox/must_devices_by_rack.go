package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/device"
)

func (c *Client) MustDevicesByRack(i int32) []*device.Device {
	result, e := c.DevicesByRack(i)
	errors.PanicOnError(e)

	return result
}
