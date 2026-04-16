package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/cluster"
	"github.com/funtimecoding/go-library/pkg/netbox/cluster_type"
	"github.com/funtimecoding/go-library/pkg/netbox/site"
	upstream "github.com/netbox-community/go-netbox/v4"
)

func (c *Client) CreateCluster(
	name string,
	t *cluster_type.Type,
	s *site.Site,
) *cluster.Cluster {
	q := upstream.NewWritableClusterRequest(
		name,
		upstream.BriefClusterTypeRequestAsClusterRequestType(
			upstream.NewBriefClusterTypeRequest(t.Name, t.Raw.Slug),
		),
	)
	q.SetScopeType("dcim.site")
	q.SetScopeId(s.Identifier)
	q.SetStatus(upstream.CLUSTERSTATUSVALUE_ACTIVE)
	result, r, e := c.client.VirtualizationAPI.VirtualizationClustersCreate(
		c.context,
	).WritableClusterRequest(*q).Execute()
	errors.PanicOnWebError(r, e)

	return cluster.New(result)
}
