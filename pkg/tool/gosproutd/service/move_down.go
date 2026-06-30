package service

func (s *Service) MoveDown(identifier uint) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.store.MoveDown(identifier)
	s.notifier.Notify()
}
