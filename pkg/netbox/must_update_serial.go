package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/device"
)

func (c *Client) MustUpdateSerial(device string, serial string) *device.Device {
	result, e := c.UpdateSerial(device, serial)
	errors.PanicOnError(e)

	return result
}
