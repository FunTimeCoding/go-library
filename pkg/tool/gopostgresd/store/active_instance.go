package store

func (s *Store) ActiveInstance(sessionIdentifier string) (string, bool) {
	v, okay := s.sessions.Load(sessionIdentifier)

	if !okay {
		return "", false
	}

	return v.(string), true
}
