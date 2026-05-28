package service

import (
	"github.com/funtimecoding/go-library/pkg/generative/ollama"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/rerank"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/store"
)

type Service struct {
	store    *store.Store
	ollama   *ollama.Client
	reranker *rerank.Reranker
}
