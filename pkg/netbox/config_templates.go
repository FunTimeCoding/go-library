package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/config_template"
)

func (c *Client) ConfigTemplates() []*config_template.Template {
	result, r, e := c.client.ExtrasAPI.ExtrasConfigTemplatesList(
		c.context,
	).Execute()
	errors.PanicOnWebError(r, e)

	return config_template.NewSlice(result.Results)
}
