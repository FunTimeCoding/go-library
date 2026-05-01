package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/device"
)

func (c *Client) MustRemoveTag(
	deviceName string,
	tag string,
) *device.Device {
	result, e := c.RemoveTag(deviceName, tag)
	errors.PanicOnError(e)

	return result
}
