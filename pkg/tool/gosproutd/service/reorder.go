package service

func (s *Service) Reorder(identifiers []uint) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.store.Reorder(identifiers)
	s.notifier.Notify()
}
