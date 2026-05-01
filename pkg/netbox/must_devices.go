package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/device"
)

func (c *Client) MustDevices() []*device.Device {
	result, e := c.Devices()
	errors.PanicOnError(e)

	return result
}
