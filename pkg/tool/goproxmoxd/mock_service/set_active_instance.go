package mock_service

func (s *Service) SetActiveInstance(
	sessionIdentifier string,
	instance string,
) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.activeInstance[sessionIdentifier] = instance
}
