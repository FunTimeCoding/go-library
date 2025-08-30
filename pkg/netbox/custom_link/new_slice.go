package custom_link

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.CustomLink) []*Link {
	var result []*Link

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
