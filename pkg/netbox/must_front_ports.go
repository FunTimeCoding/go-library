package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/front_port"
)

func (c *Client) MustFrontPorts() []*front_port.Port {
	result, e := c.FrontPorts()
	errors.PanicOnError(e)

	return result
}
