package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/custom_field"
)

func (c *Client) MustCustomFields() []*custom_field.Field {
	result, e := c.CustomFields()
	errors.PanicOnError(e)

	return result
}
