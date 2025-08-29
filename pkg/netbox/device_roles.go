package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/device_role"
)

func (c *Client) DeviceRoles() []*device_role.Role {
	if len(c.cache.DeviceRoles) != 0 {
		return c.cache.DeviceRoles
	}

	result, _, e := c.client.DcimAPI.DcimDeviceRolesList(c.context).Execute()
	errors.PanicOnError(e)
	c.cache.DeviceRoles = device_role.NewSlice(result.Results)

	return c.cache.DeviceRoles
}
