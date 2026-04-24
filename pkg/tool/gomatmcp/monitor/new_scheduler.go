package monitor

import "github.com/robfig/cron/v3"

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
