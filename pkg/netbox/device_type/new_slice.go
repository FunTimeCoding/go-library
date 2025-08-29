package device_type

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.DeviceType) []*DeviceType {
	var result []*DeviceType

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
