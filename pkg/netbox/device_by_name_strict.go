package netbox

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/netbox/device"
)

func (c *Client) DeviceByNameStrict(n string) (*device.Device, error) {
	result, e := c.DevicesByName(n)

	if e != nil {
		return nil, e
	}

	if o := len(result); o > 1 {
		for _, r := range result {
			if r.Name == n {
				return r, nil
			}
		}
	}

	if len(result) != 1 {
		return nil, fmt.Errorf(
			"expected 1 device named %s, got %d",
			n,
			len(result),
		)
	}

	return result[0], nil
}
