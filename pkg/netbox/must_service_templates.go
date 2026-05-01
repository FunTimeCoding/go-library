package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/service_template"
)

func (c *Client) MustServiceTemplates() []*service_template.Template {
	result, e := c.ServiceTemplates()
	errors.PanicOnError(e)

	return result
}
