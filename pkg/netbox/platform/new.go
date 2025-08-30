package platform

import "github.com/netbox-community/go-netbox/v4"

func New(v *netbox.Platform) *Platform {
	return &Platform{Identifier: v.GetId(), Name: v.GetName(), Raw: v}
}
