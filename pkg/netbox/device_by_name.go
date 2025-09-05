package netbox

import "github.com/funtimecoding/go-library/pkg/netbox/device"

func (c *Client) DeviceByName(n string) *device.Device {
	result := c.DevicesByName(n)

	if o := len(result); o == 1 {
		return result[0]
	} else if o > 1 {
		for _, r := range result {
			if r.Name == n {
				return r
			}
		}
	}

	return nil
}
