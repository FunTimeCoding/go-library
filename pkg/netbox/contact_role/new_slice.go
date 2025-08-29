package contact_role

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.ContactRole) []*Role {
	var result []*Role

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
