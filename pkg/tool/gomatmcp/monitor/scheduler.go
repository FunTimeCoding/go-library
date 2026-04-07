package monitor

import (
	"github.com/robfig/cron/v3"
	"sync"
)

type scheduler struct {
	cron     *cron.Cron
	schedule string
	taskFn   func() error
	entryID  cron.EntryID
	mu       sync.Mutex
	running  bool
}

func newScheduler(
	schedule string,
	taskFn func() error,
) *scheduler {
	return &scheduler{
		schedule: schedule,
		taskFn:   taskFn,
		cron:     cron.New(),
	}
}
