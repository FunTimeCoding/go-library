package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/device"
)

func (c *Client) MustUnsetSerial(device string) *device.Device {
	result, e := c.UnsetSerial(device)
	errors.PanicOnError(e)

	return result
}
