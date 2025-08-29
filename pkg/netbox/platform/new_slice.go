package platform

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.Platform) []*Platform {
	var result []*Platform

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
