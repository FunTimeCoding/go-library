package store

func (s *Store) notify() {
	s.notifier.Notify()
}
