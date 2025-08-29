package rear_port

import "github.com/netbox-community/go-netbox/v4"

func New(d *netbox.RearPort) *Port {
	return &Port{Identifier: d.GetId(), Name: d.GetName(), Raw: d}
}
