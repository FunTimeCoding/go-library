package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/contact"
)

func (c *Client) Contacts() []*contact.Contact {
	result, _, e := c.client.TenancyAPI.TenancyContactsList(
		c.context,
	).Execute()
	errors.PanicOnError(e)

	return contact.NewSlice(result.Results)
}
