package model_context

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/generative/ollama"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/rerank"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/service"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/store"
	"github.com/mark3labs/mcp-go/server"
)

type Server struct {
	server   *server.MCPServer
	service  *service.Service
	store    *store.Store
	ollama   *ollama.Client
	reranker *rerank.Reranker
	reporter face.Reporter
}
