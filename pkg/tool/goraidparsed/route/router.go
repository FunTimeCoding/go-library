package route

import (
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/getsentry/sentry-go"
)

type Router struct {
	parserPath   string
	templatePath string
	outputPath   string
	logger       *logger.Logger
	hub          *sentry.Hub
}
