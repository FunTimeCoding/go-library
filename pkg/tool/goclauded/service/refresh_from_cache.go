package service

func (s *Service) RefreshFromCache(identifier string) {
	s.statesMu.Lock()
	state, exists := s.states[identifier]
	s.statesMu.Unlock()

	if !exists {
		return
	}

	s.RefreshSession(identifier, state)
}
