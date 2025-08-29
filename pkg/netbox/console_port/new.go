package console_port

import "github.com/netbox-community/go-netbox/v4"

func New(d *netbox.ConsolePort) *Port {
	return &Port{Identifier: d.GetId(), Name: d.GetName(), Raw: d}
}
