package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/device_role"
)

func (c *Client) DeviceRoleByName(n string) *device_role.DeviceRole {
	for _, element := range c.DeviceRoles() {
		if element.Name == n {
			return element
		}
	}

	errors.NotFound(n)

	return nil
}
