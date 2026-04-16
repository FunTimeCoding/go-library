package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/virtual_machine"
	upstream "github.com/netbox-community/go-netbox/v4"
)

func (c *Client) RemoveVirtualTag(
	name string,
	tag string,
) *virtual_machine.Machine {
	vm := c.VirtualMachineByName(name)
	vm.RemoveTag(tag)
	q := upstream.NewPatchedWritableVirtualMachineWithConfigContextRequest()
	q.SetName(vm.Name)
	q.SetTags(c.tagsNestedRequest(vm.Tags))
	result, r, e := c.client.VirtualizationAPI.VirtualizationVirtualMachinesPartialUpdate(
		c.context,
		vm.Identifier,
	).PatchedWritableVirtualMachineWithConfigContextRequest(*q).Execute()
	errors.PanicOnWebError(r, e)

	return virtual_machine.New(result)
}
