package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/contact_group"
)

func (c *Client) ContactGroups() []*contact_group.Group {
	result, _, e := c.client.TenancyAPI.TenancyContactGroupsList(
		c.context,
	).Execute()
	errors.PanicOnError(e)

	return contact_group.NewSlice(result.Results)
}
