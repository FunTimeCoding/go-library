package inventory_item_role

import "github.com/netbox-community/go-netbox/v4"

func New(d *netbox.InventoryItemRole) *Role {
	return &Role{Identifier: d.GetId(), Name: d.GetName(), Raw: d}
}
