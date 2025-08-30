package module_type_profile

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.ModuleTypeProfile) []*Profile {
	var result []*Profile

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
