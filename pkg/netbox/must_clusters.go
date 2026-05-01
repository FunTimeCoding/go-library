package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/cluster"
)

func (c *Client) MustClusters() []*cluster.Cluster {
	result, e := c.Clusters()
	errors.PanicOnError(e)

	return result
}
