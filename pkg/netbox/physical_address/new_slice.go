package physical_address

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.MACAddress) []*Address {
	var result []*Address

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
