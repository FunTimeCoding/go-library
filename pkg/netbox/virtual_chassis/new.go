package virtual_chassis

import "github.com/netbox-community/go-netbox/v4"

func New(d *netbox.VirtualChassis) *Chassis {
	return &Chassis{Identifier: d.GetId(), Name: d.GetName(), Raw: d}
}
