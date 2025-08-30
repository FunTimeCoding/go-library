package module_type

import "github.com/netbox-community/go-netbox/v4"

func New(v *netbox.ModuleType) *Type {
	return &Type{Identifier: v.GetId(), Model: v.GetModel(), Raw: v}
}
