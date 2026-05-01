package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/module_bay"
)

func (c *Client) MustDeviceModuleBays(device string) []*module_bay.Bay {
	result, e := c.DeviceModuleBays(device)
	errors.PanicOnError(e)

	return result
}
