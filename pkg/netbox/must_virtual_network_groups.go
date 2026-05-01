package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/virtual_network_group"
)

func (c *Client) MustVirtualNetworkGroups() []*virtual_network_group.Group {
	result, e := c.VirtualNetworkGroups()
	errors.PanicOnError(e)

	return result
}
