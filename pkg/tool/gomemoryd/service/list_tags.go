package service

import "github.com/funtimecoding/go-library/pkg/tool/gomemoryd/store"

func (s *Service) ListTags() ([]store.TagCount, error) {
	return s.store.ListTags()
}
