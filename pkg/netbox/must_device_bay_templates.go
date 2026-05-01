package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/device_bay_template"
)

func (c *Client) MustDeviceBayTemplates() []*device_bay_template.Template {
	result, e := c.DeviceBayTemplates()
	errors.PanicOnError(e)

	return result
}
