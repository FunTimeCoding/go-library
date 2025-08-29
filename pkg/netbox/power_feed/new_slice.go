package power_feed

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.PowerFeed) []*Feed {
	var result []*Feed

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
