package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/inventory_item_role"
)

func (c *Client) MustInventoryItemRoles() []*inventory_item_role.Role {
	result, e := c.InventoryItemRoles()
	errors.PanicOnError(e)

	return result
}
