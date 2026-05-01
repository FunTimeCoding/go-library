package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/contact_group"
)

func (c *Client) MustContactGroups() []*contact_group.Group {
	result, e := c.ContactGroups()
	errors.PanicOnError(e)

	return result
}
