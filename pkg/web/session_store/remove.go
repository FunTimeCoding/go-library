package session_store

func (s *Store) Remove(v string) {
	s.Lock()
	delete(s.sessions, v)
	s.Unlock()
}
