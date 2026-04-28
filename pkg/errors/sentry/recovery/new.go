package recovery

import (
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/getsentry/sentry-go"
)

func New(
	l *logger.Logger,
	h *sentry.Hub,
) *Recovery {
	return &Recovery{logger: l, hub: h}
}
