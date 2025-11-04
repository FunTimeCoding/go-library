package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/funtimecoding/go-library/pkg/netbox/inventory_item_role"
)

func (c *Client) InventoryItemRoles() []*inventory_item_role.Role {
	result, r, e := c.client.DcimAPI.DcimInventoryItemRolesList(
		c.context,
	).Limit(constant.PageLimit).Execute()
	errors.PanicOnWebError(r, e)

	return inventory_item_role.NewSlice(result.Results)
}
