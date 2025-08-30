package virtual_chassis

import "github.com/netbox-community/go-netbox/v4"

func New(v *netbox.VirtualChassis) *Chassis {
	return &Chassis{Identifier: v.GetId(), Name: v.GetName(), Raw: v}
}
