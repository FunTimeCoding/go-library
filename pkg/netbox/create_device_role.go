package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/device_role"
	upstream "github.com/netbox-community/go-netbox/v4"
)

func (c *Client) CreateDeviceRole(name string) *device_role.Role {
	q := upstream.NewWritableDeviceRoleRequest(
		name,
		slug(name),
	)
	result, r, e := c.client.DcimAPI.DcimDeviceRolesCreate(
		c.context,
	).WritableDeviceRoleRequest(*q).Execute()
	errors.PanicOnWebError(r, e)
	c.cache.DeviceRoles = nil

	return device_role.New(result)
}
