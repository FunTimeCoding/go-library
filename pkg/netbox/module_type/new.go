package module_type

import "github.com/netbox-community/go-netbox/v4"

func New(d *netbox.ModuleType) *Type {
	return &Type{Identifier: d.GetId(), Model: d.GetModel(), Raw: d}
}
