package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/virtual_machine"
)

func (c *Client) MustVirtualMachineByName(n string) *virtual_machine.Machine {
	result, e := c.VirtualMachineByName(n)
	errors.PanicOnError(e)

	return result
}
