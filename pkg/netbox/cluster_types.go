package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/cluster_type"
)

func (c *Client) ClusterTypes() []*cluster_type.Type {
	result, _, e := c.client.VirtualizationAPI.VirtualizationClusterTypesList(
		c.context,
	).Execute()
	errors.PanicOnError(e)

	return cluster_type.NewSlice(result.Results)
}
