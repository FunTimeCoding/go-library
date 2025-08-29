package device_role

import "github.com/netbox-community/go-netbox/v4"

func New(d *netbox.DeviceRole) *DeviceRole {
	return &DeviceRole{Identifier: d.GetId(), Name: d.GetName(), Raw: d}
}
