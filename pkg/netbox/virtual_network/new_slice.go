package virtual_network

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.VLAN) []*Network {
	var result []*Network

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
