package model_context

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/service"
	"github.com/mark3labs/mcp-go/server"
)

type Server struct {
	server    *server.MCPServer
	service   *service.Service
	indexer   face.Indexer
	reporter  face.Reporter
	logger    *logger.Logger
	telemetry face.Recorder
}
