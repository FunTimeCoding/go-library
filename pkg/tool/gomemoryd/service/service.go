package service

import (
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/store"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/face"
)

type Service struct {
	store    *store.Store
	indexer  face.Indexer
	searcher face.Searcher
	lister   face.Lister
}
