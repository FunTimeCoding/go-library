package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/custom_field_choice"
)

func (c *Client) CustomFieldChoices() []*custom_field_choice.Choice {
	result, _, e := c.client.ExtrasAPI.ExtrasCustomFieldChoiceSetsList(
		c.context,
	).Execute()
	errors.PanicOnError(e)

	return custom_field_choice.NewSlice(result.Results)
}
