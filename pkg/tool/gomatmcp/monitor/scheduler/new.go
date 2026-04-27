package scheduler

import "github.com/robfig/cron/v3"

func New(
	schedule string,
	taskFunction func() error,
) *Scheduler {
	return &Scheduler{
		schedule: schedule,
		taskFunction:   taskFunction,
		cron:     cron.New(),
	}
}
