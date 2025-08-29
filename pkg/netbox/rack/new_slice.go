package rack

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.Rack) []*Rack {
	var result []*Rack

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
