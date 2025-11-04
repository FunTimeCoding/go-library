package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/funtimecoding/go-library/pkg/netbox/user"
)

func (c *Client) Users() []*user.User {
	result, r, e := c.client.UsersAPI.UsersUsersList(
		c.context,
	).Limit(constant.PageLimit).Execute()
	errors.PanicOnWebError(r, e)

	return user.NewSlice(result.Results)
}
