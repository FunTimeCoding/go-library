package module_bay

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.ModuleBay) []*Bay {
	var result []*Bay

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
