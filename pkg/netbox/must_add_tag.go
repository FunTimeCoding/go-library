package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/device"
)

func (c *Client) MustAddTag(deviceName string, tag string) *device.Device {
	result, e := c.AddTag(deviceName, tag)
	errors.PanicOnError(e)

	return result
}
