package service

import (
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/store"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/face"
)

func New(
	s *store.Store,
	i face.Indexer,
	r face.Searcher,
	l face.Lister,
) *Service {
	return &Service{store: s, indexer: i, searcher: r, lister: l}
}
