package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/virtual_machine"
)

func (c *Client) MustRemoveVirtualTag(name string, tag string) *virtual_machine.Machine {
	result, e := c.RemoveVirtualTag(name, tag)
	errors.PanicOnError(e)

	return result
}
