package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/user_group"
)

func (c *Client) MustUserGroups() []*user_group.Group {
	result, e := c.UserGroups()
	errors.PanicOnError(e)

	return result
}
