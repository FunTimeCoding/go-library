package store

func (s *Store) SetActiveInstance(
	sessionID string,
	instance string,
) {
	s.sessions.Store(sessionID, instance)
}
