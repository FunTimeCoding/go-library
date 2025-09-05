package netbox

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) DeleteInternetAddress(identifier int32) {
	result, e := c.client.IpamAPI.IpamIpAddressesDestroy(
		c.context,
		identifier,
	).Execute()
	fmt.Printf("Delete internet address: %+v\n", result)
	errors.PanicOnError(e)
}
