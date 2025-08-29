package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/device"
	"log"
)

func (c *Client) DevicesByCluster(s string) []*device.Device {
	clusters := c.ClustersByName(s)

	if len(clusters) != 1 {
		log.Panicf("unexpected: %d", len(clusters))
	}

	result, _, e := c.client.DcimAPI.DcimDevicesList(
		c.context,
	).ClusterId([]*int32{&clusters[0].Identifier}).Execute()
	errors.PanicOnError(e)

	return device.NewSlice(result.Results)
}
