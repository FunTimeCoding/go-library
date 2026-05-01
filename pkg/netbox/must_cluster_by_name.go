package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/cluster"
)

func (c *Client) MustClusterByName(n string) *cluster.Cluster {
	result, e := c.ClusterByName(n)
	errors.PanicOnError(e)

	return result
}
