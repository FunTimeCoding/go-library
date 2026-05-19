package store

func (s *Store) SetActiveInstance(
	sessionIdentifier string,
	instance string,
) {
	s.sessions.Store(sessionIdentifier, instance)
}
