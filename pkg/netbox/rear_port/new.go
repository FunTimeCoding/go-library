package rear_port

import "github.com/netbox-community/go-netbox/v4"

func New(v *netbox.RearPort) *Port {
	return &Port{Identifier: v.GetId(), Name: v.GetName(), Raw: v}
}
