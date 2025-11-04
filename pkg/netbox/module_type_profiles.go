package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/module_type_profile"
)

func (c *Client) ModuleTypeProfiles() []*module_type_profile.Profile {
	result, r, e := c.client.DcimAPI.DcimModuleTypeProfilesList(
		c.context,
	).Execute()
	errors.PanicOnWebError(r, e)

	return module_type_profile.NewSlice(result.Results)
}
