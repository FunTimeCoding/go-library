package platform

import "github.com/netbox-community/go-netbox/v4"

func New(d *netbox.Platform) *Platform {
	return &Platform{Identifier: d.GetId(), Name: d.GetName(), Raw: d}
}
