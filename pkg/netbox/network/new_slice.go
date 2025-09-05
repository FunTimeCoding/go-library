package network

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.Interface) []*Interface {
	var result []*Interface

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
