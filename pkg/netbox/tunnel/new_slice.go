package tunnel

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.Tunnel) []*Tunnel {
	var result []*Tunnel

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
