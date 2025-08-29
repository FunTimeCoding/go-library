package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/funtimecoding/go-library/pkg/netbox/rack_role"
)

func (c *Client) RackRoles() []*rack_role.Role {
	result, _, e := c.client.DcimAPI.DcimRackRolesList(
		c.context,
	).Limit(constant.PageLimit).Execute()
	errors.PanicOnError(e)

	return rack_role.NewSlice(result.Results)
}
