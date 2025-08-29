package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/cluster_group"
)

func (c *Client) ClusterGroups() []*cluster_group.Group {
	result, _, e := c.client.VirtualizationAPI.VirtualizationClusterGroupsList(
		c.context,
	).Execute()
	errors.PanicOnError(e)

	return cluster_group.NewSlice(result.Results)
}
