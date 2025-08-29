package power_port

import "github.com/netbox-community/go-netbox/v4"

func New(d *netbox.PowerPort) *Port {
	return &Port{Identifier: d.GetId(), Name: d.GetName(), Raw: d}
}
