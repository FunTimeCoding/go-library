package store

func (s *Store) Stop() {
	close(s.stop)
}
