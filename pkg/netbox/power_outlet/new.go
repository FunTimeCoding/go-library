package power_outlet

import "github.com/netbox-community/go-netbox/v4"

func New(v *netbox.PowerOutlet) *Outlet {
	return &Outlet{Identifier: v.GetId(), Name: v.GetName(), Raw: v}
}
