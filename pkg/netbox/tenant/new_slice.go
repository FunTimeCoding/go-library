package tenant

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.Tenant) []*Tenant {
	var result []*Tenant

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
