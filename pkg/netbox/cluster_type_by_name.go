package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/cluster_type"
)

func (c *Client) ClusterTypeByName(n string) *cluster_type.Type {
	for _, t := range c.ClusterTypes() {
		if t.Name == n {
			return t
		}
	}

	errors.NotFound(n)

	return nil
}
