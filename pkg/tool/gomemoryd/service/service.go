package service

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/store"
)

type Service struct {
	store    *store.Store
	indexer  face.Indexer
	searcher face.Searcher
	lister   face.Lister
}
