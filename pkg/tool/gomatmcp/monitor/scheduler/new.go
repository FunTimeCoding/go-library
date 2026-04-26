package scheduler

import "github.com/robfig/cron/v3"

func New(
	schedule string,
	taskFn func() error,
) *Scheduler {
	return &Scheduler{
		schedule: schedule,
		taskFn:   taskFn,
		cron:     cron.New(),
	}
}
