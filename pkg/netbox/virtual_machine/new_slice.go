package virtual_machine

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.VirtualMachineWithConfigContext) []*Machine {
	var result []*Machine

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
