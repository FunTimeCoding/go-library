package device_type

import "github.com/netbox-community/go-netbox/v4"

func New(d *netbox.DeviceType) *DeviceType {
	return &DeviceType{Identifier: d.GetId(), Model: d.GetModel(), Raw: d}
}
