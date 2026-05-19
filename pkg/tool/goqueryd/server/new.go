package server

import (
	"github.com/funtimecoding/go-library/pkg/generative/ollama"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/rerank"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/service"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/store"
)

func New(
	s *store.Store,
	o *ollama.Client,
	re *rerank.Reranker,
) *Server {
	return &Server{
		service:  service.New(s, o, re),
		store:    s,
		ollama:   o,
		reranker: re,
	}
}
