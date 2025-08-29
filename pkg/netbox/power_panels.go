package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/funtimecoding/go-library/pkg/netbox/power_panel"
)

func (c *Client) PowerPanels() []*power_panel.Panel {
	result, _, e := c.client.DcimAPI.DcimPowerPanelsList(
		c.context,
	).Limit(constant.PageLimit).Execute()
	errors.PanicOnError(e)

	return power_panel.NewSlice(result.Results)
}
