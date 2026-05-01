package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/power_panel"
)

func (c *Client) MustPowerPanels() []*power_panel.Panel {
	result, e := c.PowerPanels()
	errors.PanicOnError(e)

	return result
}
