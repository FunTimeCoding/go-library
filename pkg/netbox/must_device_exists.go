package netbox

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) MustDeviceExists(name string) bool {
	result, e := c.DeviceExists(name)
	errors.PanicOnError(e)

	return result
}
