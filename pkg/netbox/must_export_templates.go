package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/export_template"
)

func (c *Client) MustExportTemplates() []*export_template.Template {
	result, e := c.ExportTemplates()
	errors.PanicOnError(e)

	return result
}
