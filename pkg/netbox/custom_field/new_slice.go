package custom_field

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.CustomField) []*Field {
	var result []*Field

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
