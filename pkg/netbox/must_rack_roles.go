package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/rack_role"
)

func (c *Client) MustRackRoles() []*rack_role.Role {
	result, e := c.RackRoles()
	errors.PanicOnError(e)

	return result
}
