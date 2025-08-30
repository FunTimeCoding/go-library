package wireless_network

import "github.com/netbox-community/go-netbox/v4"

func New(v *netbox.WirelessLAN) *Network {
	return &Network{Identifier: v.GetId(), Name: v.GetDisplay(), Raw: v}
}
