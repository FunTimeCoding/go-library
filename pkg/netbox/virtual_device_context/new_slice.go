package virtual_device_context

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.VirtualDeviceContext) []*Context {
	var result []*Context

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
