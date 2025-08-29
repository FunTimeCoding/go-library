package cluster_type

import "github.com/netbox-community/go-netbox/v4"

func New(d *netbox.ClusterType) *Type {
	return &Type{Identifier: d.GetId(), Name: d.GetName(), Raw: d}
}
