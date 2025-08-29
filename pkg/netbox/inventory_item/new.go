package inventory_item

import "github.com/netbox-community/go-netbox/v4"

func New(d *netbox.InventoryItem) *Item {
	return &Item{Identifier: d.GetId(), Name: d.GetName(), Raw: d}
}
