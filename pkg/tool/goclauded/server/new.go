package server

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/service"
)

func New(
	s *service.Service,
	l *logger.Logger,
	r face.Reporter,
	harborPath string,
	sessionExportPath string,
) *Server {
	return &Server{
		service:           s,
		logger:            l,
		reporter:          r,
		harborPath:        harborPath,
		sessionExportPath: sessionExportPath,
	}
}
