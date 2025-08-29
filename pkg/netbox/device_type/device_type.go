package device_type

import "github.com/netbox-community/go-netbox/v4"

type DeviceType struct {
	Identifier int32
	Model      string

	Raw *netbox.DeviceType
}
