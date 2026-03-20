package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/configuration_template"
)

func (c *Client) ConfigurationTemplates() []*configuration_template.Template {
	result, r, e := c.client.ExtrasAPI.ExtrasConfigTemplatesList(
		c.context,
	).Execute()
	errors.PanicOnWebError(r, e)

	return configuration_template.NewSlice(result.Results)
}
