package cluster_type

import "github.com/netbox-community/go-netbox/v4"

type Type struct {
	Identifier int32
	Name       string

	Raw *netbox.ClusterType
}
