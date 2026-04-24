package monitor

import (
	"github.com/funtimecoding/go-library/pkg/chat/mattermost"
	"github.com/mattermost/mattermost/server/public/model"
)

func New(
	c *mattermost.Client,
	o *Configuration,
) *Monitor {
	m := &Monitor{
		client:               c,
		configuration:        o,
		channelCache:         make(map[string]*model.Channel),
		lastCheckMillisecond: make(map[string]int64),
	}
	m.scheduler = newScheduler(o.Schedule, m.run)

	return m
}
