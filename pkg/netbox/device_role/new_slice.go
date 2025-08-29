package device_role

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.DeviceRole) []*DeviceRole {
	var result []*DeviceRole

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
