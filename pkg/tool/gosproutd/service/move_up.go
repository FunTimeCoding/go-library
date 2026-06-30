package service

func (s *Service) MoveUp(identifier uint) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.store.MoveUp(identifier)
	s.notifier.Notify()
}
