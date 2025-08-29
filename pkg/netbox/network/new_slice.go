package network

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(
	v []netbox.Interface,
	interfaceTypes []netbox.InterfaceTypeValue,
) []*Interface {
	var result []*Interface

	for _, e := range v {
		result = append(result, New(&e, interfaceTypes))
	}

	return result
}
