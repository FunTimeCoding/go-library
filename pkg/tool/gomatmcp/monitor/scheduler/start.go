package scheduler

import "github.com/funtimecoding/go-library/pkg/errors"

func (s *Scheduler) Start() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.running {
		panic("scheduler already running")
	}

	entry, e := s.cron.AddFunc(
		s.schedule,
		func() {
			defer func() {
				if v := recover(); v != nil {
					if s.hub != nil {
						s.hub.Recover(v)
					}

					s.logger.Plain("scheduler recovered from panic: %v", v)
				}
			}()
			s.task()
		},
	)
	errors.PanicOnError(e)
	s.entry = entry
	s.cron.Start()
	s.running = true
}
