package wireless_network_group

import "github.com/netbox-community/go-netbox/v4"

func New(v *netbox.WirelessLANGroup) *Group {
	return &Group{Identifier: v.GetId(), Name: v.GetName(), Raw: v}
}
