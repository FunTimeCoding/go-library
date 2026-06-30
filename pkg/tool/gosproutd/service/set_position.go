package service

func (s *Service) SetPosition(
	identifier uint,
	target int,
) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.store.SetPosition(identifier, target)
	s.notifier.Notify()
}
