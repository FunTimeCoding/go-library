package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/cluster"
)

func (c *Client) ClustersByName(s string) []*cluster.Cluster {
	result, _, e := c.client.VirtualizationAPI.VirtualizationClustersList(
		c.context,
	).NameIc([]string{s}).Execute()
	errors.PanicOnError(e)

	return cluster.NewSlice(result.Results)
}
