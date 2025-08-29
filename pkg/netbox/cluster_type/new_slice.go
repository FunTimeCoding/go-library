package cluster_type

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.ClusterType) []*Type {
	var result []*Type

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
