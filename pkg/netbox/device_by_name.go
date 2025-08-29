package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors/unexpected"
	"github.com/funtimecoding/go-library/pkg/netbox/device"
)

func (c *Client) DeviceByName(
	n string,
	panic bool,
) *device.Device {
	result := c.DevicesByName(n)

	if panic {
		unexpected.Count(1, len(result))
	}

	if len(result) > 0 {
		return result[0]
	}

	return nil
}
