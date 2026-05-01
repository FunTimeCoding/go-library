package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/cluster_type"
)

func (c *Client) MustClusterTypeByName(n string) *cluster_type.Type {
	result, e := c.ClusterTypeByName(n)
	errors.PanicOnError(e)

	return result
}
