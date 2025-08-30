package inventory_item

import "github.com/netbox-community/go-netbox/v4"

func New(v *netbox.InventoryItem) *Item {
	return &Item{Identifier: v.GetId(), Name: v.GetName(), Raw: v}
}
