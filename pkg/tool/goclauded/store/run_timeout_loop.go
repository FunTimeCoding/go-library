package store

import "time"

func (s *Store) RunTimeoutLoop() {
	for {
		time.Sleep(5 * time.Minute)
		s.RunTimeoutSweep()
	}
}
