package virtual_chassis

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.VirtualChassis) []*Chassis {
	var result []*Chassis

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
