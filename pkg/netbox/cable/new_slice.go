package cable

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.Cable) []*Cable {
	var result []*Cable

	for _, element := range v {
		result = append(result, New(&element))
	}

	return result
}
