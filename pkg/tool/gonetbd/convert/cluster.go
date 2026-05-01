package convert

import (
	"github.com/funtimecoding/go-library/pkg/netbox/cluster"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/generated/server"
)

func Cluster(c *cluster.Cluster) server.Cluster {
	result := server.Cluster{Identifier: c.Identifier, Name: c.Name}

	if c.Raw.Type.Name != "" {
		result.Type = &c.Raw.Type.Name
	}

	return result
}
