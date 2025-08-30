package virtual_network

import "github.com/netbox-community/go-netbox/v4"

func New(v *netbox.VLAN) *Network {
	return &Network{Identifier: v.GetId(), Name: v.GetName(), Raw: v}
}
