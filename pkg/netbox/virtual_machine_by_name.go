package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/virtual_machine"
)

func (c *Client) VirtualMachineByName(n string) *virtual_machine.Machine {
	for _, m := range c.VirtualMachines() {
		if m.Name == n {
			return m
		}
	}

	errors.NotFound(n)

	return nil
}
