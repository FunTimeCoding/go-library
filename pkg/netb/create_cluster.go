package netb

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/client"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) CreateCluster(
	name string,
	clusterType string,
	site string,
) string {
	result, e := c.client.CreateCluster(
		c.context,
		client.CreateClusterRequest{
			Name: name,
			Type: clusterType,
			Site: site,
		},
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
