package service

import "github.com/funtimecoding/go-library/pkg/tool/gotelemetryd/store"

func (s *Service) Query(o *store.QueryOption) ([]store.UsageEvent, error) {
	return s.store.Recent(o)
}
