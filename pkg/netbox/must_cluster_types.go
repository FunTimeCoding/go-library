package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/cluster_type"
)

func (c *Client) MustClusterTypes() []*cluster_type.Type {
	result, e := c.ClusterTypes()
	errors.PanicOnError(e)

	return result
}
