package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/device_role"
)

func (c *Client) DeviceRoleByName(n string) *device_role.Role {
	for _, r := range c.DeviceRoles() {
		if r.Name == n {
			return r
		}
	}

	errors.NotFound(n)

	return nil
}
