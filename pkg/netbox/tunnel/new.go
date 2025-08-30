package tunnel

import "github.com/netbox-community/go-netbox/v4"

func New(v *netbox.Tunnel) *Tunnel {
	return &Tunnel{Identifier: v.GetId(), Name: v.GetName(), Raw: v}
}
