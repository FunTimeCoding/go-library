package contact_group

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.ContactGroup) []*Group {
	var result []*Group

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
