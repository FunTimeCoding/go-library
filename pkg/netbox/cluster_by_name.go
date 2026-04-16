package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/cluster"
)

func (c *Client) ClusterByName(n string) *cluster.Cluster {
	for _, cl := range c.ClustersByName(n) {
		if cl.Name == n {
			return cl
		}
	}

	errors.NotFound(n)

	return nil
}
