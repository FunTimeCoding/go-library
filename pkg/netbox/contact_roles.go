package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/contact_role"
)

func (c *Client) ContactRoles() []*contact_role.Role {
	result, r, e := c.client.TenancyAPI.TenancyContactRolesList(
		c.context,
	).Execute()
	errors.PanicOnWebError(r, e)

	return contact_role.NewSlice(result.Results)
}
