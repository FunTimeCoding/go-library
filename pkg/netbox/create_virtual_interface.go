package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/virtual_machine"
	upstream "github.com/netbox-community/go-netbox/v4"
)

func (c *Client) CreateVirtualInterface(
	vm *virtual_machine.Machine,
	name string,
) *upstream.VMInterface {
	q := upstream.NewWritableVMInterfaceRequest(
		upstream.BriefVirtualMachineRequestAsPatchedVirtualDiskRequestVirtualMachine(
			upstream.NewBriefVirtualMachineRequest(vm.Name),
		),
		name,
	)
	result, r, e := c.client.VirtualizationAPI.VirtualizationInterfacesCreate(
		c.context,
	).WritableVMInterfaceRequest(*q).Execute()
	errors.PanicOnWebError(r, e)

	return result
}
