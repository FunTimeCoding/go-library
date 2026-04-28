package monitor

import (
	"github.com/funtimecoding/go-library/pkg/chat/mattermost"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/tool/gomatmcp/monitor/scheduler"
	"github.com/getsentry/sentry-go"
	"github.com/mattermost/mattermost/server/public/model"
)

func New(
	c *mattermost.Client,
	o *Configuration,
	l *logger.Logger,
	h *sentry.Hub,
) *Monitor {
	m := &Monitor{
		client:               c,
		configuration:        o,
		logger:               l,
		hub:                  h,
		channelCache:         make(map[string]*model.Channel),
		lastCheckMillisecond: make(map[string]int64),
	}
	m.scheduler = scheduler.New(o.Schedule, m.run, l, h)

	return m
}
