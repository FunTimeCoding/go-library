package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/cluster_type"
)

func (c *Client) MustCreateClusterType(name string) *cluster_type.Type {
	result, e := c.CreateClusterType(name)
	errors.PanicOnError(e)

	return result
}
