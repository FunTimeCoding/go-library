package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/virtual_machine"
	upstream "github.com/netbox-community/go-netbox/v4"
	"log"
)

func (c *Client) VirtualMachineInterfaceByName(
	vm *virtual_machine.Machine,
	name string,
) *upstream.VMInterface {
	result, r, e := c.client.VirtualizationAPI.VirtualizationInterfacesList(
		c.context,
	).VirtualMachineId([]int32{vm.Identifier}).Execute()
	errors.PanicOnWebError(r, e)

	for _, i := range result.Results {
		if i.GetName() == name {
			return &i
		}
	}

	log.Panicf("interface %s not found for VM %s", name, vm.Name)

	return nil
}
