package config_context

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.ConfigContext) []*Context {
	var result []*Context

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
