package service

func (s *Service) RunTimeoutSweep() {
	s.sweepInactivity()
	s.sweepCompleteTimeout()
}
