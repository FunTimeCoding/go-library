package netbox

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) DeleteDevice(identifier int32) {
	result, e := c.client.DcimAPI.DcimDevicesDestroy(
		c.context,
		identifier,
	).Execute()
	fmt.Printf("Delete device result: %+v\n", result)
	errors.PanicOnError(e)
}
