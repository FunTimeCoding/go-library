package netbox

import (
	"github.com/funtimecoding/go-library/pkg/netbox/virtual_machine"
	upstream "github.com/netbox-community/go-netbox/v4"
)

func (c *Client) CreateVirtualInterface(
	vm *virtual_machine.Machine,
	name string,
) (*upstream.VMInterface, error) {
	q := upstream.NewWritableVMInterfaceRequest(
		upstream.BriefVirtualMachineRequestAsPatchedVirtualDiskRequestVirtualMachine(
			upstream.NewBriefVirtualMachineRequest(vm.Name),
		),
		name,
	)
	result, _, e := c.client.VirtualizationAPI.VirtualizationInterfacesCreate(
		c.context,
	).WritableVMInterfaceRequest(*q).Execute()

	return result, e
}
