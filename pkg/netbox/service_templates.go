package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/funtimecoding/go-library/pkg/netbox/service_template"
)

func (c *Client) ServiceTemplates() []*service_template.Template {
	result, _, e := c.client.IpamAPI.IpamServiceTemplatesList(
		c.context,
	).Limit(constant.PageLimit).Execute()
	errors.PanicOnError(e)

	return service_template.NewSlice(result.Results)
}
