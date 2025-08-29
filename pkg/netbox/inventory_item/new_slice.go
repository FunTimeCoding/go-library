package inventory_item

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.InventoryItem) []*Item {
	var result []*Item

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
