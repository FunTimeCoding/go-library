package virtual_machine

import "github.com/netbox-community/go-netbox/v4"

type Machine struct {
	Identifier int32
	Name       string

	Raw *netbox.VirtualMachineWithConfigContext
}
