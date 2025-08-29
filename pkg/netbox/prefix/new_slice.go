package prefix

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.Prefix) []*Prefix {
	var result []*Prefix

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
