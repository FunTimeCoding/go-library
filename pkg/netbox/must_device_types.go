package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/device_type"
)

func (c *Client) MustDeviceTypes() []*device_type.Type {
	result, e := c.DeviceTypes()
	errors.PanicOnError(e)

	return result
}
