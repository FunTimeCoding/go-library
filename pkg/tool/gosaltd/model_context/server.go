package model_context

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/tool/gosaltd/runner"
	"github.com/funtimecoding/go-library/pkg/tool/gosaltd/store"
	"github.com/mark3labs/mcp-go/server"
)

type Server struct {
	server   *server.MCPServer
	runner   *runner.Runner
	store    *store.Store
	reporter face.Reporter
}
