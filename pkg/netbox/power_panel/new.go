package power_panel

import "github.com/netbox-community/go-netbox/v4"

func New(d *netbox.PowerPanel) *Panel {
	return &Panel{Identifier: d.GetId(), Name: d.GetName(), Raw: d}
}
