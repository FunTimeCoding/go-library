package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors/unexpected"
	"github.com/funtimecoding/go-library/pkg/netbox/device"
)

func (c *Client) DeviceByNameStrict(n string) *device.Device {
	result := c.DevicesByName(n)

	if o := len(result); o > 1 {
		for _, r := range result {
			if r.Name == n {
				return r
			}
		}
	}

	unexpected.Count(1, len(result))

	return result[0]
}
