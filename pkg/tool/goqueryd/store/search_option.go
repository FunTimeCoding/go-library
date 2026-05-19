package store

import "github.com/funtimecoding/go-library/pkg/tool/goqueryd/rerank"

type SearchOption struct {
	Query      string
	Limit      int
	Collection string
	Full       bool
	SourceType string
	Reranker   *rerank.Reranker
}
