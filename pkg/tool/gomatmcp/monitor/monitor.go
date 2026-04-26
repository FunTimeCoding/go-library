package monitor

import (
	"github.com/funtimecoding/go-library/pkg/chat/mattermost"
	"github.com/funtimecoding/go-library/pkg/tool/gomatmcp/monitor/scheduler"
	"github.com/mattermost/mattermost/server/public/model"
	"sync"
)

type Monitor struct {
	client               *mattermost.Client
	configuration        *Configuration
	scheduler            *scheduler.Scheduler
	channelCache         map[string]*model.Channel
	lastCheckMillisecond map[string]int64
	notifyChannel        *model.Channel
	username             string
	running              bool
	mutex                sync.RWMutex
}
