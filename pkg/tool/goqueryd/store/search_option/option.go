package search_option

import "github.com/funtimecoding/go-library/pkg/tool/goqueryd/rerank"

type Option struct {
	Query      string
	Limit      int
	Collection string
	Full       bool
	Mode       string
	Metadata   map[string]string
	Reranker   *rerank.Reranker
	Exclude    []string
}
