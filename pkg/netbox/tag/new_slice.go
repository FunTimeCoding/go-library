package tag

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.Tag) []*Tag {
	var result []*Tag

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
