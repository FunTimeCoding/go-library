package store

func (s *Store) Close() {
	s.client.Close()
}
