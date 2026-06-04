package store

import "github.com/funtimecoding/go-library/pkg/tool/goqueryd/rerank"

type SearchOption struct {
	Query      string
	Limit      int
	Collection string
	Full       bool
	Metadata   map[string]string
	Reranker   *rerank.Reranker
}
