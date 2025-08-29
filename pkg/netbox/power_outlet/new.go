package power_outlet

import "github.com/netbox-community/go-netbox/v4"

func New(d *netbox.PowerOutlet) *Outlet {
	return &Outlet{Identifier: d.GetId(), Name: d.GetName(), Raw: d}
}
