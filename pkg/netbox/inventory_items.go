package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/funtimecoding/go-library/pkg/netbox/inventory_item"
)

func (c *Client) InventoryItems() []*inventory_item.Item {
	result, r, e := c.client.DcimAPI.DcimInventoryItemsList(
		c.context,
	).Limit(constant.PageLimit).Execute()
	errors.PanicOnWebError(r, e)

	return inventory_item.NewSlice(result.Results)
}
