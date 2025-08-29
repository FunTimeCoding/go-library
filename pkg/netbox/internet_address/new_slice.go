package internet_address

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.IPAddress) []*Address {
	var result []*Address

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
