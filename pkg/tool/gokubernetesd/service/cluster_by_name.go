package service

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/tool/gokubernetesd/service/cluster"
)

func (s *Service) ClusterByName(name string) (*cluster.Cluster, error) {
	c, okay := s.clusters[name]

	if !okay {
		return nil, fmt.Errorf("unknown cluster: %s", name)
	}

	return c, nil
}
