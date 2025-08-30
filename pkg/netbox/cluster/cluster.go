package cluster

import "github.com/netbox-community/go-netbox/v4"

type Cluster struct {
	Identifier int32
	Name       string

	Raw *netbox.Cluster
}
