package rack_type

import "github.com/netbox-community/go-netbox/v4"

func New(v *netbox.RackType) *Type {
	return &Type{Identifier: v.GetId(), Name: v.GetModel(), Raw: v}
}
