package cluster

import "github.com/netbox-community/go-netbox/v4"

func New(v *netbox.Cluster) *Cluster {
	return &Cluster{Identifier: v.GetId(), Name: v.GetName(), Raw: v}
}
