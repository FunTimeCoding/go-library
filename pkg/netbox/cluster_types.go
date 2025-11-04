package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/cluster_type"
)

func (c *Client) ClusterTypes() []*cluster_type.Type {
	result, r, e := c.client.VirtualizationAPI.VirtualizationClusterTypesList(
		c.context,
	).Execute()
	errors.PanicOnWebError(r, e)

	return cluster_type.NewSlice(result.Results)
}
