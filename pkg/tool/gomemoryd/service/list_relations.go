package service

import "github.com/funtimecoding/go-library/pkg/tool/gomemoryd/store"

func (s *Service) ListRelations(identifier int64) ([]store.Relation, error) {
	return s.store.ListRelations(identifier)
}
