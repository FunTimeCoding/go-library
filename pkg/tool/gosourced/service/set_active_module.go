package service

func (s *Service) SetActiveModule(
	sessionIdentifier string,
	module string,
) {
	s.sessions.Store(sessionIdentifier, module)
}
