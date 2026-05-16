package convert

import (
	"github.com/funtimecoding/go-library/pkg/netbox/virtual_machine"
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/generated/server"
)

func VirtualMachines(v []*virtual_machine.Machine) []*server.VirtualMachine {
	result := make([]*server.VirtualMachine, 0, len(v))

	for _, m := range v {
		result = append(result, VirtualMachine(m))
	}

	return result
}
