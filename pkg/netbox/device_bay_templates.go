package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/device_bay_template"
)

func (c *Client) DeviceBayTemplates() []*device_bay_template.Template {
	result, _, e := c.client.DcimAPI.DcimDeviceBayTemplatesList(
		c.context,
	).Execute()
	errors.PanicOnError(e)

	return device_bay_template.NewSlice(result.Results)
}
