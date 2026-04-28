package recovery

import (
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/getsentry/sentry-go"
)

type Recovery struct {
	logger *logger.Logger
	hub    *sentry.Hub
}
