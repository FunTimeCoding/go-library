package contact

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.Contact) []*Contact {
	var result []*Contact

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
