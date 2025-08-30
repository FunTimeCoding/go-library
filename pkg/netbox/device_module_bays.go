package netbox

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/module_bay"
)

func (c *Client) DeviceModuleBays(device string) []*module_bay.Bay {
	fmt.Printf("Name: %s\n", device)
	result, _, e := c.client.DcimAPI.DcimModuleBaysList(
		c.context,
	).Device([]*string{&device}).Execute()
	errors.PanicOnError(e)

	return module_bay.NewSlice(result.Results)
}
