package device_role

import "github.com/netbox-community/go-netbox/v4"

func New(d *netbox.DeviceRole) *Role {
	return &Role{Identifier: d.GetId(), Name: d.GetName(), Raw: d}
}
