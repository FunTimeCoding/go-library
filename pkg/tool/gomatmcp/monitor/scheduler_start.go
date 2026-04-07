package monitor

import (
	"fmt"
	"log"
)

func (s *scheduler) start() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.running {
		return fmt.Errorf("scheduler already running")
	}

	id, e := s.cron.AddFunc(
		s.schedule,
		func() {
			if f := s.taskFn(); f != nil {
				log.Printf("scheduled task error: %v", f)
			}
		},
	)

	if e != nil {
		return fmt.Errorf(
			"invalid cron schedule %q: %w",
			s.schedule,
			e,
		)
	}

	s.entryID = id
	s.cron.Start()
	s.running = true

	return nil
}
