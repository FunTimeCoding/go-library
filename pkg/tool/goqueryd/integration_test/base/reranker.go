package base

import "github.com/funtimecoding/go-library/pkg/tool/goqueryd/rerank"

func (s *Server) Reranker() *rerank.Reranker {
	return s.reranker
}
