package service

import "github.com/funtimecoding/go-library/pkg/tool/goqueryd/store/index"

func (s *Service) Index(collection string) *index.Index {
	return s.store.Index(collection)
}
