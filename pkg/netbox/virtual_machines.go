package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/funtimecoding/go-library/pkg/netbox/virtual_machine"
)

func (c *Client) VirtualMachines() []*virtual_machine.Machine {
	result, _, e := c.client.VirtualizationAPI.VirtualizationVirtualMachinesList(
		c.context,
	).Limit(constant.PageLimit).Execute()
	errors.PanicOnError(e)

	return virtual_machine.NewSlice(result.Results)
}
