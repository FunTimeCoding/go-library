package netbox

import (
	"github.com/funtimecoding/go-library/pkg/netbox/cluster_type"
	upstream "github.com/netbox-community/go-netbox/v4"
)

func (c *Client) CreateClusterType(name string) (*cluster_type.Type, error) {
	q := upstream.NewClusterTypeRequest(
		name,
		slug(name),
	)
	result, _, e := c.client.VirtualizationAPI.VirtualizationClusterTypesCreate(
		c.context,
	).ClusterTypeRequest(*q).Execute()

	if e != nil {
		return nil, e
	}

	return cluster_type.New(result), nil
}
