package service

import "time"

func (s *Service) RunUsageLoop() {
	s.pollUsage()

	for {
		time.Sleep(time.Minute)
		s.pollUsage()
	}
}
