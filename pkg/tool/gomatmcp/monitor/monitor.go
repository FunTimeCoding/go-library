package monitor

import (
	"github.com/funtimecoding/go-library/pkg/chat/mattermost"
	"github.com/mattermost/mattermost/server/public/model"
	"sync"
)

type Monitor struct {
	client               *mattermost.Client
	configuration        *Configuration
	scheduler            *scheduler
	mutex                sync.RWMutex
	channelCache         map[string]*model.Channel
	lastCheckMillisecond map[string]int64
	notifyChannel        *model.Channel
	username             string
	running              bool
}

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
