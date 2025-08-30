package virtual_network_group

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.VLANGroup) []*Group {
	var result []*Group

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
