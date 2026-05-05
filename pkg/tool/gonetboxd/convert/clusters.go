package convert

import (
	"github.com/funtimecoding/go-library/pkg/netbox/cluster"
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/generated/server"
)

func Clusters(v []*cluster.Cluster) []server.Cluster {
	result := make([]server.Cluster, 0, len(v))

	for _, c := range v {
		result = append(result, Cluster(c))
	}

	return result
}
