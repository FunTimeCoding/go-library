package virtual_machine

import "github.com/netbox-community/go-netbox/v4"

func New(v *netbox.VirtualMachineWithConfigContext) *Machine {
	return &Machine{Identifier: v.GetId(), Name: v.GetName(), Raw: v}
}
