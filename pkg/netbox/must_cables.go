package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/cable"
)

func (c *Client) MustCables() []*cable.Cable {
	result, e := c.Cables()
	errors.PanicOnError(e)

	return result
}
