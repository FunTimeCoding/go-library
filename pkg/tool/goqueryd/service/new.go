package service

import (
	"github.com/funtimecoding/go-library/pkg/generative/ollama"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/rerank"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/store"
)

func New(
	s *store.Store,
	o *ollama.Client,
	re *rerank.Reranker,
) *Service {
	return &Service{store: s, ollama: o, reranker: re}
}
