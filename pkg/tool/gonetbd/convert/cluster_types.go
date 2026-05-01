package convert

import (
	"github.com/funtimecoding/go-library/pkg/netbox/cluster_type"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/generated/server"
)

func ClusterTypes(v []*cluster_type.Type) []server.ClusterType {
	result := make([]server.ClusterType, 0, len(v))

	for _, t := range v {
		result = append(result, ClusterType(t))
	}

	return result
}
