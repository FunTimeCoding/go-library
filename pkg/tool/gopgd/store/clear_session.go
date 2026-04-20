package store

func (s *Store) ClearSession(sessionID string) {
	s.sessions.Delete(sessionID)
}
