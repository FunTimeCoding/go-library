package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/cluster"
)

func (c *Client) Clusters() []*cluster.Cluster {
	result, r, e := c.client.VirtualizationAPI.VirtualizationClustersList(
		c.context,
	).Execute()
	errors.PanicOnWebError(r, e)

	return cluster.NewSlice(result.Results)
}
