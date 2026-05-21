package service

func (s *Service) ReleaseByCallsign(c string) {
	s.store.ReleaseByCallsign(c)
	s.notify()
}
