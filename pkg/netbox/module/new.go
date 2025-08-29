package module

import "github.com/netbox-community/go-netbox/v4"

func New(v *netbox.Module) *Module {
	return &Module{
		Identifier:  v.GetId(),
		Name:        v.GetDisplay(),
		Description: v.GetDescription(),
		Raw:         v,
	}
}
