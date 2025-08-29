package module_bay

import "github.com/netbox-community/go-netbox/v4"

func New(b *netbox.ModuleBay) *Bay {
	return &Bay{
		Identifier:  b.GetId(),
		Name:        b.GetName(),
		Description: b.GetDescription(),
		Raw:         b,
	}
}
