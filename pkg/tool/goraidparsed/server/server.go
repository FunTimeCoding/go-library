package server

import (
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/getsentry/sentry-go"
)

type Server struct {
	parserPath   string
	templatePath string
	outputPath   string
	logger       *logger.Logger
	hub          *sentry.Hub
}
