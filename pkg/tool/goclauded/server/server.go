package server

import (
	"github.com/funtimecoding/go-library/pkg/generative/anthropic/claude"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/service"
)

type Server struct {
	service           *service.Service
	claude            *claude.Client
	logger            *logger.Logger
	harborPath        string
	sessionExportPath string
}
