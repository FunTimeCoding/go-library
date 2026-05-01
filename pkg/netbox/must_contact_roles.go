package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/contact_role"
)

func (c *Client) MustContactRoles() []*contact_role.Role {
	result, e := c.ContactRoles()
	errors.PanicOnError(e)

	return result
}
