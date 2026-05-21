package service

import "time"

func (s *Service) RunTimeoutLoop() {
	for {
		time.Sleep(5 * time.Minute)
		s.RunTimeoutSweep()
	}
}
