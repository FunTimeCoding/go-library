package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/custom_field_choice"
)

func (c *Client) MustCustomFieldChoices() []*custom_field_choice.Choice {
	result, e := c.CustomFieldChoices()
	errors.PanicOnError(e)

	return result
}
