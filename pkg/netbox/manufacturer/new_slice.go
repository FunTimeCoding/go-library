package manufacturer

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.Manufacturer) []*Manufacturer {
	var result []*Manufacturer

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
