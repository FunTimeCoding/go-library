package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/funtimecoding/go-library/pkg/netbox/virtual_machine"
)

func (c *Client) VirtualMachines() []*virtual_machine.Machine {
	result, r, e := c.client.VirtualizationAPI.VirtualizationVirtualMachinesList(
		c.context,
	).Limit(constant.PageLimit).Execute()
	errors.PanicOnWebError(r, e)

	return virtual_machine.NewSlice(result.Results)
}
