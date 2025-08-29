package site

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.Site) []*Site {
	var result []*Site

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
