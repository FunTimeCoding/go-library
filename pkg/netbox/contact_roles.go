package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/contact_role"
)

func (c *Client) ContactRoles() []*contact_role.Role {
	result, _, e := c.client.TenancyAPI.TenancyContactRolesList(
		c.context,
	).Execute()
	errors.PanicOnError(e)

	return contact_role.NewSlice(result.Results)
}
