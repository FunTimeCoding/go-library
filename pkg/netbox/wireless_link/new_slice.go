package wireless_link

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.WirelessLink) []*Link {
	var result []*Link

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
