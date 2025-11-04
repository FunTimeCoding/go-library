package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/cluster"
)

func (c *Client) ClustersByName(s string) []*cluster.Cluster {
	result, r, e := c.client.VirtualizationAPI.VirtualizationClustersList(
		c.context,
	).NameIc([]string{s}).Execute()
	errors.PanicOnWebError(r, e)

	return cluster.NewSlice(result.Results)
}
