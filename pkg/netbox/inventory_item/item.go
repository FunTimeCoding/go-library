package inventory_item

import "github.com/netbox-community/go-netbox/v4"

type Item struct {
	Identifier int32
	Name       string

	Raw *netbox.InventoryItem
}
