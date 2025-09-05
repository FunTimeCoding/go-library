package cable

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.Cable) []*Cable {
	var result []*Cable

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
