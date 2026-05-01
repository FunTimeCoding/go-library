package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/device_role"
)

func (c *Client) MustDeviceRoleByName(n string) *device_role.Role {
	result, e := c.DeviceRoleByName(n)
	errors.PanicOnError(e)

	return result
}
