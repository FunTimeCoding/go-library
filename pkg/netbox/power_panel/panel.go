package power_panel

import "github.com/netbox-community/go-netbox/v4"

type Panel struct {
	Identifier int32
	Name       string

	Raw *netbox.PowerPanel
}
