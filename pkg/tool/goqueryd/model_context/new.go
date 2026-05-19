package model_context

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/generative/ollama"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/rerank"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/service"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/store"
	"github.com/mark3labs/mcp-go/server"
)

func New(
	s *store.Store,
	o *ollama.Client,
	re *rerank.Reranker,
	r face.Reporter,
	version string,
) *Server {
	result := &Server{
		server: server.NewMCPServer(
			constant.Identity.Name(),
			version,
			server.WithToolCapabilities(true),
			server.WithInstructions(constant.Identity.Instructions()),
		),
		service:  service.New(s, o, re),
		store:    s,
		ollama:   o,
		reranker: re,
		reporter: r,
	}
	result.register()
	result.registerResources()

	return result
}
