package service

func (s *Service) BackfillSessions() {
	for _, e := range s.store.AllSessions(0, 0) {
		s.EnrichSession(e.Identifier)
	}
}
