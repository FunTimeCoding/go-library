package service

import "github.com/funtimecoding/go-library/pkg/tool/gomemoryd/store"

func (s *Service) GetMemoryHistory(identifier int64) ([]store.Version, error) {
	return s.store.GetMemoryHistory(identifier)
}
