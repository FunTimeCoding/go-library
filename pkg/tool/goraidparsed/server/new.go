package server

import (
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/getsentry/sentry-go"
)

func New(
	parserPath string,
	templatePath string,
	outputPath string,
	l *logger.Logger,
	h *sentry.Hub,
) *Server {
	return &Server{
		parserPath:   parserPath,
		templatePath: templatePath,
		outputPath:   outputPath,
		logger:       l,
		hub:          h,
	}
}
