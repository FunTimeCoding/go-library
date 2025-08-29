package cluster

import "github.com/netbox-community/go-netbox/v4"

func New(d *netbox.Cluster) *Cluster {
	return &Cluster{Identifier: d.GetId(), Name: d.GetName(), Raw: d}
}
