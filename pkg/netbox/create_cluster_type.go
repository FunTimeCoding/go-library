package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/cluster_type"
	upstream "github.com/netbox-community/go-netbox/v4"
)

func (c *Client) CreateClusterType(name string) *cluster_type.Type {
	q := upstream.NewClusterTypeRequest(
		name,
		slug(name),
	)
	result, r, e := c.client.VirtualizationAPI.VirtualizationClusterTypesCreate(
		c.context,
	).ClusterTypeRequest(*q).Execute()
	errors.PanicOnWebError(r, e)

	return cluster_type.New(result)
}
