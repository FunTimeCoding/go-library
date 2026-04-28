package scheduler

import (
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/getsentry/sentry-go"
	"github.com/robfig/cron/v3"
)

func New(
	schedule string,
	task func(),
	l *logger.Logger,
	h *sentry.Hub,
) *Scheduler {
	return &Scheduler{
		schedule: schedule,
		task:     task,
		logger:   l,
		hub:      h,
		cron:     cron.New(),
	}
}
