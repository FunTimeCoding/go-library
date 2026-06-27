package mock_service

func (s *Service) ActiveInstance(sessionIdentifier string) (string, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	v, okay := s.activeInstance[sessionIdentifier]

	return v, okay
}
