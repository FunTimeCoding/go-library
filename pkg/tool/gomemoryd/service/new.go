package service

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/store"
)

func New(
	s *store.Store,
	i face.Indexer,
) *Service {
	return &Service{Store: s, indexer: i}
}
