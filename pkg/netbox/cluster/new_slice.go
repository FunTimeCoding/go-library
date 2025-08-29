package cluster

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.Cluster) []*Cluster {
	var result []*Cluster

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
