package power_panel

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.PowerPanel) []*Panel {
	var result []*Panel

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
