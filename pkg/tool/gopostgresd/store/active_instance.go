package store

func (s *Store) ActiveInstance(sessionID string) (string, bool) {
	v, okay := s.sessions.Load(sessionID)

	if !okay {
		return "", false
	}

	return v.(string), true
}
