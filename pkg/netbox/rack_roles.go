package netbox

import (
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/funtimecoding/go-library/pkg/netbox/rack_role"
)

func (c *Client) RackRoles() ([]*rack_role.Role, error) {
	result, _, e := c.client.DcimAPI.DcimRackRolesList(
		c.context,
	).Limit(constant.PageLimit).Execute()

	if e != nil {
		return nil, e
	}

	return rack_role.NewSlice(result.Results), nil
}
