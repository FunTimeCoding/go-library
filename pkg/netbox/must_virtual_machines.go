package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/virtual_machine"
)

func (c *Client) MustVirtualMachines() []*virtual_machine.Machine {
	result, e := c.VirtualMachines()
	errors.PanicOnError(e)

	return result
}
