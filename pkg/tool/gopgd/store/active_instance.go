package store

func (s *Store) ActiveInstance(sessionID string) (string, bool) {
	v, ok := s.sessions.Load(sessionID)

	if !ok {
		return "", false
	}

	return v.(string), true
}
