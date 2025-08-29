package inventory_item_role

import "github.com/netbox-community/go-netbox/v4"

type Role struct {
	Identifier int32
	Name       string

	Raw *netbox.InventoryItemRole
}
