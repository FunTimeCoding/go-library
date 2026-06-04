package service

import "github.com/funtimecoding/go-library/pkg/tool/goqueryd/store"

func (s *Service) ListSourceTypes() []store.SourceTypeTag {
	return s.store.ListSourceTypes()
}
