package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/contact_group"
)

func (c *Client) ContactGroups() []*contact_group.Group {
	result, r, e := c.client.TenancyAPI.TenancyContactGroupsList(
		c.context,
	).Execute()
	errors.PanicOnWebError(r, e)

	return contact_group.NewSlice(result.Results)
}
