package service

import "github.com/funtimecoding/go-library/pkg/tool/gokubernetesd/store"

func New(s *store.Store) *Service {
	return &Service{
		clusters: buildClusters(),
		store:    s,
	}
}
