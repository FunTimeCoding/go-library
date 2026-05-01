package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/contact"
)

func (c *Client) MustContacts() []*contact.Contact {
	result, e := c.Contacts()
	errors.PanicOnError(e)

	return result
}
