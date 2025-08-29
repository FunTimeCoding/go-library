package virtual_machine

import "github.com/netbox-community/go-netbox/v4"

func New(d *netbox.VirtualMachineWithConfigContext) *Machine {
	return &Machine{Identifier: d.GetId(), Name: d.GetName(), Raw: d}
}
