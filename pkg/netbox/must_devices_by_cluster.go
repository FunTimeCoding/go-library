package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/device"
)

func (c *Client) MustDevicesByCluster(s string) []*device.Device {
	result, e := c.DevicesByCluster(s)
	errors.PanicOnError(e)

	return result
}
