package device_role

import "github.com/netbox-community/go-netbox/v4"

type DeviceRole struct {
	Identifier int32
	Name       string

	Raw *netbox.DeviceRole
}
