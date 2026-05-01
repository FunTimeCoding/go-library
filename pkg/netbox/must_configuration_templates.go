package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/configuration_template"
)

func (c *Client) MustConfigurationTemplates() []*configuration_template.Template {
	result, e := c.ConfigurationTemplates()
	errors.PanicOnError(e)

	return result
}
