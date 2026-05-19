package store

func (s *Store) ClearSession(sessionIdentifier string) {
	s.sessions.Delete(sessionIdentifier)
}
