package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/tunnel_group"
)

func (c *Client) MustTunnelGroups() []*tunnel_group.Group {
	result, e := c.TunnelGroups()
	errors.PanicOnError(e)

	return result
}
