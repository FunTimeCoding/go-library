package zone

import "github.com/hetznercloud/hcloud-go/v2/hcloud"

func NewSlice(v []*hcloud.Zone) []*Zone {
	var result []*Zone

	for _, e := range v {
		result = append(result, New(e))
	}

	return result
}
