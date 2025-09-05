package netbox

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) DeletePhysicalAddress(identifier int32) {
	result, e := c.client.DcimAPI.DcimMacAddressesDestroy(
		c.context,
		identifier,
	).Execute()
	fmt.Printf("Delete physical address: %+v\n", result)
	errors.PanicOnError(e)
}
