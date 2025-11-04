package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/cluster_group"
)

func (c *Client) ClusterGroups() []*cluster_group.Group {
	result, r, e := c.client.VirtualizationAPI.VirtualizationClusterGroupsList(
		c.context,
	).Execute()
	errors.PanicOnWebError(r, e)

	return cluster_group.NewSlice(result.Results)
}
