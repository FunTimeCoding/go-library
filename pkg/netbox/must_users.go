package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/user"
)

func (c *Client) MustUsers() []*user.User {
	result, e := c.Users()
	errors.PanicOnError(e)

	return result
}
