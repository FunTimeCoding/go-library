package power_outlet

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.PowerOutlet) []*Outlet {
	var result []*Outlet

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
