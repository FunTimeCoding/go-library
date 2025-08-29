package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/unexpected"
	"github.com/funtimecoding/go-library/pkg/netbox/physical_address"
	"net"
)

func (c *Client) PhysicalAddress(
	a net.HardwareAddr,
	panic bool,
) *physical_address.Address {
	result, _, e := c.client.DcimAPI.DcimMacAddressesList(
		c.context,
	).MacAddressNic([]string{a.String()}).Execute()
	errors.PanicOnError(e)

	if panic {
		unexpected.Count(1, len(result.Results))
	}

	if len(result.Results) > 0 {
		return physical_address.New(&result.Results[0])
	}

	return nil
}
