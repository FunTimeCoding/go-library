package server

import (
	"github.com/funtimecoding/go-library/pkg/generative/ollama"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/rerank"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/service"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/store"
)

type Server struct {
	service  *service.Service
	store    *store.Store
	ollama   *ollama.Client
	reranker *rerank.Reranker
}
