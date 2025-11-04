package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/funtimecoding/go-library/pkg/netbox/user_group"
)

func (c *Client) UserGroups() []*user_group.Group {
	result, r, e := c.client.UsersAPI.UsersGroupsList(
		c.context,
	).Limit(constant.PageLimit).Execute()
	errors.PanicOnWebError(r, e)

	return user_group.NewSlice(result.Results)
}
