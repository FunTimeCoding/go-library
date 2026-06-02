package service

func (s *Service) MoveUp(identifier uint) {
	s.store.MoveUp(identifier)
	s.notifier.Notify()
}
