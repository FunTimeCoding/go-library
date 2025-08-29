package device

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.DeviceWithConfigContext) []*Device {
	var result []*Device

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
