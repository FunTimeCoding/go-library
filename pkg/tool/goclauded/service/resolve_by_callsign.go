package service

func (s *Service) ResolveByCallsign(c string) string {
	return s.store.ResolveByCallsign(c)
}
