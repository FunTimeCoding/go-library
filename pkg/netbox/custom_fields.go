package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/custom_field"
)

func (c *Client) CustomFields() []*custom_field.Field {
	result, _, e := c.client.ExtrasAPI.ExtrasCustomFieldsList(
		c.context,
	).Execute()
	errors.PanicOnError(e)

	return custom_field.NewSlice(result.Results)
}
