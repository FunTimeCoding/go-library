package rear_port

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.RearPort) []*Port {
	var result []*Port

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
