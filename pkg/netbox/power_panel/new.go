package power_panel

import "github.com/netbox-community/go-netbox/v4"

func New(v *netbox.PowerPanel) *Panel {
	return &Panel{Identifier: v.GetId(), Name: v.GetName(), Raw: v}
}
