package module

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.Module) []*Module {
	var result []*Module

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
