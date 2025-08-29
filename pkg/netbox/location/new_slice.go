package location

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.Location) []*Location {
	var result []*Location

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
