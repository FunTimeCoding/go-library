package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/export_template"
)

func (c *Client) ExportTemplates() []*export_template.Template {
	result, _, e := c.client.ExtrasAPI.ExtrasExportTemplatesList(
		c.context,
	).Execute()
	errors.PanicOnError(e)

	return export_template.NewSlice(result.Results)
}
