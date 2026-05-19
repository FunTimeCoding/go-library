package server

import (
	"github.com/funtimecoding/go-library/pkg/generative/anthropic/claude"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/service"
)

func New(
	s *service.Service,
	c *claude.Client,
	l *logger.Logger,
	harborPath string,
	sessionExportPath string,
) *Server {
	return &Server{
		service:           s,
		claude:            c,
		logger:            l,
		harborPath:        harborPath,
		sessionExportPath: sessionExportPath,
	}
}
