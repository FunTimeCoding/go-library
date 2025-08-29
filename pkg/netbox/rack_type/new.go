package rack_type

import "github.com/netbox-community/go-netbox/v4"

func New(d *netbox.RackType) *Type {
	return &Type{Identifier: d.GetId(), Name: d.GetModel(), Raw: d}
}
