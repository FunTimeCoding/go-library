package data_source

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.DataSource) []*Source {
	var result []*Source

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
