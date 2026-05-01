package netbox

import "github.com/funtimecoding/go-library/pkg/netbox/device"

func (c *Client) DeviceByName(n string) (*device.Device, error) {
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
	} else if o == 1 {
		return result[0], nil
	}

	return nil, nil
}
